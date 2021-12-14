package services

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/smokers10/go-diagem.git/domain"
)

type produkServiceImpl struct {
	produkRepository        domain.ProdukRepository
	ProdukFotoRepository    domain.ProdukFotoRepository
	ProdukVariasiRepository domain.ProdukVariasiRepository
}

func ProdukService(produk *domain.ProdukRepository, foto *domain.ProdukFotoRepository, variasi *domain.ProdukVariasiRepository) domain.ProdukService {
	return &produkServiceImpl{
		produkRepository:        *produk,
		ProdukFotoRepository:    *foto,
		ProdukVariasiRepository: *variasi,
	}
}

//Untuk Admin
func (p *produkServiceImpl) Create(req *domain.Produk) *domain.Response {
	// dekalarasi var
	res := domain.Response{}
	fotoProduk := []domain.ProdukFoto{}

	// buat uuidv4
	id, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat membuat ID produk"
		res.Status = http.StatusInternalServerError
		return &res
	}
	req.ID = id.String()

	// buat slug
	req.Slug = slug.Make(req.Nama)

	// panggil repository utk menyimpan produk
	produk, err := p.produkRepository.Create(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat menyimpan produk"
		res.Status = http.StatusInternalServerError
		return &res
	}

	// menyimpan foto produk dari type64 ke file citra
	path := fmt.Sprintf("public/uploads/produk/%s", req.ID)
	if len(req.ProdukFoto) > 0 {
		if err := os.MkdirAll(path, 0755); err != nil {
			panic(err)
		}

		for index, file := range req.ProdukFoto {
			dec, err := base64.StdEncoding.DecodeString(file.Base64)
			if err != nil {
				panic(err)
			}

			img := fmt.Sprintf("foto-produk-%d.%s", index+1, file.Format)
			f, err := os.Create(filepath.Join(path, filepath.Base(img)))
			if err != nil {
				panic(err)
			}
			defer f.Close()

			if _, err := f.Write(dec); err != nil {
				panic(err)
			}

			if err := f.Sync(); err != nil {
				panic(err)
			}

			// panggil repository untuk menyimpan data foto
			var isUtama bool
			if index == 0 {
				isUtama = true
			} else {
				isUtama = false
			}

			fotodata := domain.ProdukFoto{
				ProdukID: req.ID,
				Path:     fmt.Sprintf("/uploads/produk/%s/foto-produk-%d.%s", req.ID, index+1, file.Format),
				IsUtama:  isUtama,
			}

			fotoProduk = append(fotoProduk, fotodata)
		}
	}

	// menyimpan foto produk ke database
	for _, data := range fotoProduk {
		fotoID, _ := uuid.NewRandom()
		data.ID = fotoID.String()

		err := p.ProdukFotoRepository.Create(&data)

		if err != nil {
			fmt.Println(err)
			res.Message = "error saat membuat menyimpan foto, silahkan upload ulang dihalaman edit produk"
			res.Status = http.StatusInternalServerError
			return &res
		}
	}

	// menyimpan variasi produk (jika ada)
	if req.IsHasVariant || len(req.Variasi) > 0 {
		for _, variasi := range req.Variasi {
			variasiID, _ := uuid.NewRandom()
			variasi.ID = variasiID.String()
			variasi.ProdukID = req.ID

			_, err := p.ProdukVariasiRepository.Create(&variasi)
			if err != nil {
				fmt.Println(err)
				res.Message = "error saat menyimpan variasi"
				res.Status = http.StatusInternalServerError
				return &res
			}
		}

		req.Harga = int(req.Variasi[0].Harga)
	}

	res.Data = produk
	res.Message = "produk berhasil disimpan"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (p *produkServiceImpl) Update(req *domain.Produk) *domain.Response {
	// dekalarasi var
	res := domain.Response{}

	// buat slug
	req.Slug = slug.Make(req.Nama)

	// panggil repository terkait
	produk, err := p.produkRepository.Update(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat update produk"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = produk
	res.Message = "produk berhasil diupdate"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (p *produkServiceImpl) Delete(id string) *domain.Response {
	// dekalarasi var
	res := domain.Response{}

	// panggil repository terkait
	if err := p.produkRepository.Delete(id); err != nil {
		fmt.Println(err)
		res.Message = "error saat menghapus produk"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "produk berhasil dihapus"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

//Untuk Umum
func (p *produkServiceImpl) Read(filter *domain.ProdukFilter) *domain.Response {
	// dekalarasi var
	res := domain.Response{}

	// klarifikasi filter
	clarifyFilter(filter)

	// panggil repository terkait
	produk, err := p.produkRepository.Read(filter)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil produk"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = produk
	res.Message = "produk berhasil diambil"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func clarifyFilter(filter *domain.ProdukFilter) {
	// sort by
	switch filter.ShortBy {
	case "harga-tertinggi":
		filter.ClarifyOrder.TableName = "produk.harga"
		filter.ClarifyOrder.OrderMethod = "DESC"
	case "harga-terendah":
		filter.ClarifyOrder.TableName = "produk.harga"
		filter.ClarifyOrder.OrderMethod = "ASC"
	default:
		filter.ClarifyOrder.TableName = "produk.created_at"
		filter.ClarifyOrder.OrderMethod = "DESC"
	}
}

func (p *produkServiceImpl) Detail(id string) *domain.Response {
	// dekalarasi var
	res := domain.Response{}

	// panggil repository terkait
	produk, err := p.produkRepository.ByID(id)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil produk"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = produk
	res.Message = "produk berhasil diambil"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (p *produkServiceImpl) DetailBySlug(slug string) *domain.Response {
	// dekalarasi var
	res := domain.Response{}

	// panggil repository terkait
	produk, err := p.produkRepository.BySlugs(slug)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil produk"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = produk
	res.Message = "produk berhasil diambil"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}
