package domain

type AlamatOrigin struct {
	ID            int    `json:"id,omitempty" form:"id"`
	KontakAdmin   string `json:"kontak_admin,omitempty" form:"kontak_admin"`
	AlamatLengkap string `json:"alamat_lengkap,omitempty" form:"alamat_lengkap"`
	ProvinsiID    string `json:"provinsi_id,omitempty" form:"provinsi"`
	KotaID        string `json:"kota_id,omitempty" form:"kota_id"`
	NamaProvinsi  string `json:"nama_provinsi,omitempty" form:"nama_provinsi"`
	NamaKota      string `json:"nama_kota,omitempty" form:"nama_kota"`
	KDPos         string `json:"kd_pos,omitempty" form:"kd_pos"`
}

type AlamatOriginRepository interface {
	Read() (*AlamatOrigin, error)
	Update(req *AlamatOrigin) (*AlamatOrigin, error)
}

type AlamatOriginService interface {
	Read() *Response
	Update(req *AlamatOrigin) *Response
}
