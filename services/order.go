package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/infrastructure/etc"
)

type orderService struct {
	orderRepository         domain.OrderRepository
	produkRepository        domain.ProdukRepository
	produkVariasiRepository domain.ProdukVariasiRepository
	cartRepository          domain.CartRepository
	orderItemRepository     domain.OrderItemRepository
	userRepository          domain.UserRepository
	alamatRepository        domain.AlamatRepository
	bankRepository          domain.BankRepository
	orderBayarRepository    domain.OrderBayarRepository
}

func OrderService(order *domain.OrderRepository, produk *domain.ProdukRepository, variasi *domain.ProdukVariasiRepository,
	cart *domain.CartRepository, item *domain.OrderItemRepository, user *domain.UserRepository,
	alamat *domain.AlamatRepository, bank *domain.BankRepository, bayar *domain.OrderBayarRepository) domain.OrderService {
	return &orderService{
		orderRepository:         *order,
		produkRepository:        *produk,
		produkVariasiRepository: *variasi,
		cartRepository:          *cart,
		orderItemRepository:     *item,
		userRepository:          *user,
		alamatRepository:        *alamat,
		bankRepository:          *bank,
		orderBayarRepository:    *bayar,
	}
}

func (osc *orderService) Create(req *domain.Order) *domain.Response {
	res := domain.Response{}
	itemDetails := []etc.ItemDetail{}
	total := req.Total

	// ambil data cart
	carts, err := osc.cartRepository.Read(req.UserID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil data order"
		res.Status = 500
		return &res
	}

	// ambil data user
	user, err := osc.userRepository.ByID(req.UserID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil data pengguna"
		res.Status = 500
		return &res
	}

	// ambil data alamat
	alamat, err := osc.alamatRepository.ByID(req.AlamatID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil data alamat pengguna"
		res.Status = 500
		return &res
	}

	// // ambil data virtual account yang sesuai dengan konfigurasi dari admin
	// bank, err := osc.bankRepository.Segmented()
	// if err != nil {
	// 	fmt.Println(err)
	// 	res.Message = "error saat mengambil data bank"
	// 	res.Status = 500
	// 	return &res
	// }

	// check stok pada tiap cart item
	for i := 0; i < len(carts); i++ {
		quantity := carts[i].Quantity
		if carts[i].Produk.IsHasVariant {
			variantStok := carts[i].Variasi.Stok

			if osc.IsStokSufficient(variantStok, quantity) {
				res.Message = fmt.Sprintf("Produk %s Variasi %s Sudah Habis", carts[i].Produk.Nama, carts[i].Variasi.Variant)
				res.Status = 200
				return &res
			}
		} else {
			produkStok := carts[i].Produk.Stok

			if osc.IsStokSufficient(produkStok, quantity) {
				res.Message = fmt.Sprintf("Produk %s Sudah Habis", carts[i].Produk.Nama)
				res.Status = 200
				return &res
			}
		}
	}

	// Set Item Detail
	itemDetails = append(itemDetails, osc.SetItemDetail(carts)...)

	// assign ongkir ke item detail agar gross amount midtrans tidak error
	ongkir_detail := etc.ItemDetail{
		ID:       "ongkir",
		Price:    req.Ongkir,
		Quantity: 1,
		Name:     req.Kurir + " - " + req.PaketKurir,
	}

	itemDetails = append(itemDetails, ongkir_detail)

	// generate id
	id, _ := uuid.NewRandom()
	req.ID = id.String()

	// midtrans
	midtrans := etc.Midtrans{
		TransactionDetail: etc.TransactionDetail{
			OrderID:     req.ID,
			GrossAmount: total,
		},
		CustomerDetail: etc.CustomerDetail{
			Email: user.Email,
			Phone: user.HP,
			ShippingAddress: etc.ShippingAddress{
				Phone:      alamat.Phone,
				FirstName:  alamat.Penerima,
				Address:    fmt.Sprintf("%s %s %s", alamat.Alamat, alamat.NamaKota, alamat.NamaProvinsi),
				City:       alamat.NamaKota,
				PostalCode: alamat.KDPos,
			},
		},
		ItemDetail: itemDetails,
		BCA:        etc.BCA{VANumber: "4566545676543", SubCompanyNumber: "G549190193"},
		BNI:        etc.BNI{VANumber: "123456543345"},
		BRI:        etc.BRI{VANumber: "67876545678"},
	}

	midres, err := midtrans.Transaction()
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat koneksi midtrans"
		res.Status = 500
		return &res
	}

	// generate invoice
	year, month, day := time.Now().Date()
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 10000
	randomNumb := rand.Intn(max-min+1) + min
	invoiceNo := fmt.Sprintf("DGM-%d-%d%d%d", randomNumb, day, int(month), year)
	req.InvoiceNo = invoiceNo

	// ================================ vital part! ==================================================
	// mulai transaction db
	tx, err := osc.orderRepository.GetSQLInstance().Begin()
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat memulai transaksi"
		res.Status = 500
		return &res
	}

	// buat order
	if err := osc.orderRepository.Create(req, tx); err != nil {
		fmt.Println(err)
		tx.Rollback()
		res.Message = "error saat menyimpan data order"
		res.Status = 500
		return &res
	}

	// buat pembayaran order
	orderPayment := domain.OrderBayar{
		OrderCheckoutID: req.ID,
		Jumlah:          float32(total),
		Token:           midres.Token,
		RedirectURL:     midres.RedirectURL,
	}

	if err := osc.orderBayarRepository.Create(&orderPayment, tx); err != nil {
		fmt.Println(err)
		tx.Rollback()
		res.Message = "error saat membuat pembayaran"
		res.Status = 500
		return &res
	}

	// commit / rollback
	if midres.RedirectURL == "" || midres.Token == "" {
		tx.Rollback()
		res.Message = "checkout gagal coba lagi"
		res.Status = 200
		return &res
	}

	tx.Commit()

	// pindahkan data cart ke order item
	for _, cart := range carts {
		orderItem := domain.OrderItem{
			OrderCheckoutID: id.String(),
			VariasiID:       cart.VariasiID,
			ProdukID:        cart.ProdukID,
			Quantity:        cart.Quantity,
		}

		if err := osc.orderItemRepository.Create(&orderItem); err != nil {
			fmt.Println(err)
			res.Message = "error saat menyimpan data order item"
			res.Status = 500
			return &res
		}

		// naikan jumlah penjualan
		osc.produkRepository.UpdateSellingPoint(cart.ProdukID, cart.Quantity)

		// update stok
		if !cart.Produk.IsHasVariant {
			if err := osc.produkRepository.UpdateStok(cart.Produk.ID, cart.Quantity, "subtraction"); err != nil {
				fmt.Println(err)
				res.Message = "kesalahan saat update stok produk"
				res.Status = 200
				return &res
			}
		} else {
			if err := osc.produkVariasiRepository.UpdateStokVariant(cart.VariasiID, cart.Quantity, "subtraction"); err != nil {
				fmt.Println(err)
				res.Message = "kesalahan saat update stok produk"
				res.Status = 200
				return &res
			}
		}

		osc.cartRepository.Delete(&domain.Cart{ID: cart.ID, UserID: cart.UserID})
	}
	// ================================ end vital part! =============================================

	res.Data = map[string]string{"order_id": req.ID}
	res.Message = "checkout berhasil"
	res.Status = 200
	res.Success = true
	return &res
}

func (osc *orderService) ReadAll() *domain.Response {
	res := domain.Response{}

	order, err := osc.orderRepository.Read()
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil data order"
		res.Status = 500
		return &res
	}

	// check status transkasi
	for _, o := range order {
		if o.OrderBayar.Status != "settlement" {
			midtrans, err := etc.MidtransSnap().CheckStatus(o.ID)
			if err != nil {
				fmt.Println(err)
				return &domain.Response{
					Message: "error saat check status transaksi",
					Status:  500,
				}
			}

			if midtrans.StatusCode != "404" {
				// jika ada perubahan
				if o.OrderBayar.Status != midtrans.TransactionStatus {
					err := osc.orderBayarRepository.UpdateStatus(o.OrderBayar.Token, midtrans.TransactionStatus)
					if err != nil {
						fmt.Println(err)
						return &domain.Response{
							Message: "error update status transkasi",
							Status:  500,
						}
					}
				}
			}
		}
	}

	res.Data = order
	res.Message = "data order berhasil diambil"
	res.Status = 200
	return &res
}

func (osc *orderService) Detail(orderID string) *domain.Response {
	// ambil detail order
	order, err := osc.orderRepository.GetByID(orderID)
	if err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat mengambil data order detail",
			Status:  500,
		}
	}

	// check apakah status pembayaran lunas atau belum
	// jika belum lakukan request check status transaksi midtrans
	if order.OrderBayar.Status != "settlement" {
		midtrans, err := etc.MidtransSnap().CheckStatus(orderID)
		if err != nil {
			fmt.Println(err)
			return &domain.Response{
				Message: "error saat check status transaksi",
				Status:  500,
			}
		}

		if midtrans.StatusCode != "404" {
			// jika ada perubahan
			if order.OrderBayar.Status != midtrans.TransactionStatus {
				err := osc.orderBayarRepository.UpdateStatus(order.OrderBayar.Token, midtrans.TransactionStatus)
				if err != nil {
					fmt.Println(err)
					return &domain.Response{
						Message: "error update status transkasi",
						Status:  500,
					}
				}
			}
		}
	}

	return &domain.Response{
		Message: "detail transaksi berhasil di ambil",
		Success: true,
		Data:    order,
		Status:  200,
	}
}

func (osc *orderService) UpdateSR(orderID string, status string, resi string) *domain.Response {
	res := domain.Response{}

	if err := osc.orderRepository.UpdateSR(orderID, status, resi); err != nil {
		fmt.Println(err)
		res.Message = "error saat update status order"
		res.Status = 500
		return &res
	}

	res.Message = "status order berhasil diubah"
	res.Status = 200
	return &res
}

func (osc *orderService) ReadByUser(userID int) *domain.Response {
	res := domain.Response{}

	order, err := osc.orderRepository.GetByUserID(userID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil data order"
		res.Status = 500
		return &res
	}

	// check status transkasi
	for _, o := range order {
		if o.OrderBayar.Status != "settlement" {
			midtrans, err := etc.MidtransSnap().CheckStatus(o.ID)
			if err != nil {
				fmt.Println(err)
				return &domain.Response{
					Message: "error saat check status transaksi",
					Status:  500,
				}
			}

			if midtrans.StatusCode != "404" {
				// jika ada perubahan
				if o.OrderBayar.Status != midtrans.TransactionStatus {
					err := osc.orderBayarRepository.UpdateStatus(o.OrderBayar.Token, midtrans.TransactionStatus)
					if err != nil {
						fmt.Println(err)
						return &domain.Response{
							Message: "error update status transkasi",
							Status:  500,
						}
					}
				}
			}
		}
	}

	res.Data = order
	res.Message = "data order berhasil diambil"
	res.Status = 200
	return &res
}

// OUT OF ORDER SERVICE SCOPE
func (osc *orderService) IsStokSufficient(stok int, quantity int) bool {
	check := stok - quantity
	return check < 0
}

// OUT OF ORDER SERVICE SCOPE - mendapatkan item detail dari carts
func (osc *orderService) SetItemDetail(carts []domain.CartDetail) (itemDetails []etc.ItemDetail) {
	for _, cart := range carts {
		if cart.Produk.IsHasVariant {
			if cart.Produk.Discount > 0 {
				var pengali float32 = float32(cart.Produk.Discount) / 100
				price := float32(float32(cart.Variasi.Harga) - (float32(cart.Variasi.Harga) * pengali))
				itemDetails = append(itemDetails, osc.AssignItemDetail(&cart, int(price)))
			} else {
				price := cart.Variasi.Harga
				itemDetails = append(itemDetails, osc.AssignItemDetail(&cart, int(price)))
			}
		} else {
			if cart.Produk.Discount > 0 {
				var pengali float32 = float32(cart.Produk.Discount) / 100
				price := float32(float32(cart.Produk.Harga) - (float32(cart.Produk.Harga) * pengali))
				itemDetails = append(itemDetails, osc.AssignItemDetail(&cart, int(price)))
			} else {
				price := float32(cart.Produk.Harga)
				itemDetails = append(itemDetails, osc.AssignItemDetail(&cart, int(price)))
			}
		}
	}

	return itemDetails
}

// OUT OF ORDER SERVICE SCOPE - assign item detail dari carts ke struct etc.ItemDetail
func (osc *orderService) AssignItemDetail(cart *domain.CartDetail, processedPrice int) (itemDetail etc.ItemDetail) {
	itemDetail.ID = cart.ProdukID
	itemDetail.Name = cart.Produk.Nama
	itemDetail.Quantity = cart.Quantity
	itemDetail.Price = processedPrice

	return itemDetail
}
