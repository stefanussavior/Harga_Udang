package controller

import (
	"harga_udang/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TambahHargaUdang struct {
	Tanggal       time.Time
	Nama_penambak string
	Harga         int64
	Ukuran        int32
}

type UpdateHargaUdang struct {
	Tanggal       time.Time
	Nama_penambak string
	Harga         int64
	Ukuran        int32
}

func LihatDataUdang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var DataUdang models.HargaUdang
	db.First(&DataUdang)

	c.JSON(http.StatusOK, gin.H{"Data": DataUdang})
}

func TambahDataUdang(c *gin.Context) {
	var TambahUdang models.HargaUdang

	if err := c.ShouldBindJSON(&TambahUdang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&TambahUdang)
	c.JSON(http.StatusOK, gin.H{"Data : ": &TambahUdang})
}

func CariDataUdang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var Cari models.HargaUdang

	if err := db.Where("id = ?", c.Param("id")).First(&Cari).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Data": Cari})
}

func UpdateDataUdang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var UpdateUdang models.HargaUdang

	if err := db.Where("id = ?", c.Param("id")).First(&UpdateUdang).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	var InputUpdateUdang models.HargaUdang

	if err := c.ShouldBindJSON(&InputUpdateUdang); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	db.Model(&UpdateUdang).Updates(&InputUpdateUdang)
	c.JSON(http.StatusOK, gin.H{"Data Update": &InputUpdateUdang})
}

func HapusHargaUdang(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var HapusUdang models.HargaUdang

	if err := db.Where("id = ?", c.Param("id")).First(&HapusUdang).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	db.Delete(&HapusUdang)
	c.JSON(http.StatusOK, gin.H{"Pesan : ": " Data Telah Terhapus"})
}
