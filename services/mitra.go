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
	mitra, err := m.mitraRepository.ByEmail(req.Email)
	if err != nil {
		fmt.Println(err.Error())
		res.Message = "error saat check ketersedian akun mitra"
		res.Status = http.StatusInternalServerError
		return &res
	}

	if mitra.Email != "" {
		res.Message = fmt.Sprintf("mitra dengan email %s sudah terdaftar", req.Email)
		res.Status = http.StatusConflict
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
