package services

import (
	"fmt"
	"net/http"

	"github.com/smokers10/go-diagem.git/domain"
)

type mitraServiceImpl struct {
	mitraRepository domain.MitraRepository
}

func MitraService(mitra *domain.MitraRepository) domain.MitraService {
	return &mitraServiceImpl{mitraRepository: *mitra}
}

func (m *mitraServiceImpl) Create(req *domain.Mitra) *domain.Response {
	//deklarasi var
	res := domain.Response{}

	//check email apakah sudah digunakan
	emailOK, err := m.emailIsOK(req.Email)
	if err != nil {
		fmt.Println(err.Error())
		res.Message = "error saat check email"
		res.Status = http.StatusInternalServerError
		return &res
	}

	if !emailOK {
		res.Message = "email sudah digunakan"
		return &res
	}

	// check seller id
	sellerIDOK, err := m.sellerIDIsOk(req.SellerID)
	if err != nil {
		fmt.Println(err.Error())
		res.Message = "error saat check email"
		res.Status = http.StatusInternalServerError
		return &res
	}

	if !sellerIDOK {
		res.Message = "seller ID sudah digunakan"
		return &res
	}

	//panggil repository simpan
	insertedMitra, err := m.mitraRepository.Create(req)
	if err != nil {
		fmt.Println(err.Error())
		res.Message = "error saat menyimpan data mitra"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = insertedMitra
	res.Message = "mitra baru berhasil ditambahkan"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (m *mitraServiceImpl) Update(req *domain.Mitra) *domain.Response {
	//deklarasi var
	res := domain.Response{}

	updatedMitra, err := m.mitraRepository.Update(req)
	if err != nil {
		fmt.Println(err.Error())
		res.Message = "error saat update data mitra"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = updatedMitra
	res.Message = "mitra berhasil diupdate"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (m *mitraServiceImpl) Delete(id int) *domain.Response {
	//deklarasi var
	res := domain.Response{}

	err := m.mitraRepository.Delete(id)
	if err != nil {
		fmt.Println(err.Error())
		res.Message = "error saat hapus data mitra"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "mitra berhasil dihapus"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

//Untuk Umum
func (m *mitraServiceImpl) Read() *domain.Response {
	//deklarasi var
	res := domain.Response{}

	mitra, err := m.mitraRepository.Read()
	if err != nil {
		fmt.Println(err.Error())
		res.Message = "error saat mengambil data mitra"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = mitra
	res.Message = "mitra berhasil diambil"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (m *mitraServiceImpl) GetOne(id int) *domain.Response {
	//deklarasi var
	res := domain.Response{}

	mitra, err := m.mitraRepository.ByID(id)
	if err != nil {
		fmt.Println(err.Error())
		res.Message = "error saat mengambil data mitra"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = mitra
	res.Message = "mitra berhasil diambil"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (m *mitraServiceImpl) ReadWithFilter(req *domain.Mitra) *domain.Response {
	//deklarasi var
	res := domain.Response{}

	mitra, err := m.mitraRepository.ReadWithFilter(req)
	if err != nil {
		fmt.Println(err.Error())
		res.Message = "error saat mengambil data mitra"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = mitra
	res.Message = "mitra berhasil diambil"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

// out of abstrak checking email
func (m *mitraServiceImpl) emailIsOK(email string) (bool, error) {
	mitra, err := m.mitraRepository.ByEmail(email)
	if err != nil {
		return false, err
	}

	if mitra.Email != "" {
		return false, nil
	}

	return true, nil
}

// out of abstrak checking seller id
func (m *mitraServiceImpl) sellerIDIsOk(sellerID string) (bool, error) {
	mitra, err := m.mitraRepository.BySellerID(sellerID)
	if err != nil {
		return false, err
	}

	if mitra.SellerID != "" {
		return false, nil
	}

	return true, nil
}
