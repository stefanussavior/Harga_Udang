package models

import (
	"time"

	"gorm.io/gorm"
)

type HargaUdang struct {
	gorm.Model
	Tanggal       time.Time
	Nama_penambak string
	Harga         int64
	Ukuran        int32
}

func (HargaUdang) TableName() string {
	return "harga_udang"
}
