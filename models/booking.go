package models

import (
	"time"
)

type Booking struct {
	Id        string  `gorm:"type:varchar(300); primaryKey" json:"id"`
	Ktp       string  `gorm:"type:varchar(250)" json:"ktp"`
	Bpjs      string  `gorm:"type:varchar(250)" json:"bpjs"`
	Rm		  string  `json:"rm"`
	Bayar     string  `gorm:"type:varchar(100)" json:"bayar"`
	Nama      string  `gorm:"type:varchar(250)" json:"nama"`
	Poli      string  `gorm:"type:varchar(150)" json:"poli"`
	Selesai   string  `gorm:"type:varchar(100)" json:"selesai"`
	Tanggal   int  	  `json:"tanggal"`
	Bulan     string  `gorm:"type:varchar(150)" json:"bulan"`
	Tahun     int     `json:"tahun"`
	Tempat    string  `gorm:"type:varchar(250)" json:"tempat"`
	Lahir     string  `gorm:"type:varchar(150)" json:"lahir"`
	Telepon   string  `gorm:"type:varchar(100)" json:"telepon"`
	Alamat    string  `gorm:"type:varchar(250)" json:"alamat"`
	Kelamin   string  `gorm:"type:varchar(100)" json:"kelamin"`
	Agama     string  `gorm:"type:varchar(100)" json:"agama"`
	Ibu       string  `gorm:"type:varchar(250)" json:"ibu"`
	Keluarga  string  `gorm:"type:varchar(250)" json:"keluarga"`
	Status    string  `gorm:"type:varchar(250)" json:"status"`
	Suku      string  `gorm:"type:varchar(150)" json:"suku"`
	Pekerjaan string  `gorm:"type:varchar(150)" json:"pekerjaan"`
	Propinsi  string  `gorm:"type:varchar(150)" json:"propinsi"`
	Kabupaten string  `gorm:"type:varchar(150)" json:"kabupaten"`
	Kecamatan string  `gorm:"type:varchar(150)" json:"kecamatan"`
	Desa      string  `gorm:"type:varchar(150)" json:"desa"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}