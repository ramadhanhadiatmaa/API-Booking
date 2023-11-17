package models

import (
	"time"
)

type Booking struct {
	Id        string  `gorm:"primaryKey;" json:"id"`
	Ktp       int     `json:"ktp"`
	Bpjs      int     `json:"bpjs"`
	Rm		  int     `json:"rm"`
	Bayar     string  `gorm:"type:varchar(250)" json:"bayar"`
	Nama      string  `gorm:"type:varchar(250)" json:"nama"`
	Poli      string  `gorm:"type:varchar(250)" json:"poli"`
	Tanggal   int  	  `json:"tanggal"`
	Bulan     string  `gorm:"type:varchar(250)" json:"bulan"`
	Tahun     int     `json:"tahun"`
	Tempat    string  `gorm:"type:varchar(150)" json:"tempat"`
	Lahir     string  `gorm:"type:varchar(150)" json:"lahir"`
	Telepon   string  `gorm:"type:varchar(150)" json:"telepon"`
	Alamat    string  `gorm:"type:varchar(250)" json:"alamat"`
	Kelamin   string  `gorm:"type:varchar(250)" json:"kelamin"`
	Agama     string  `gorm:"type:varchar(250)" json:"agama"`
	Ibu       string  `gorm:"type:varchar(250)" json:"ibu"`
	Keluarga  string  `gorm:"type:varchar(250)" json:"keluarga"`
	Status    string  `gorm:"type:varchar(250)" json:"status"`
	Suku      string  `gorm:"type:varchar(250)" json:"suku"`
	Pekerjaan string  `gorm:"type:varchar(250)" json:"pekerjaan"`
	Propinsi  string  `gorm:"type:varchar(250)" json:"propinsi"`
	Kabupaten string  `gorm:"type:varchar(250)" json:"kabupaten"`
	Kecamatan string  `gorm:"type:varchar(250)" json:"kecamatan"`
	Desa      string  `gorm:"type:varchar(250)" json:"desa"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}