package database

import (
	"fmt"
	"harga_udang/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	User := "root"
	Pass := ""
	Host := "localhost"
	Port := "3306"
	DbName := "db_fistx_harga_udang"

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", User, Pass, Host, Port, DbName)

	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("terkoneksi ke dalam database")
	}
	db.AutoMigrate(&models.HargaUdang{})
	return db
}
