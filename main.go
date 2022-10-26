package main

import (
	"harga_udang/controller"
	"harga_udang/database"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.ConnectDatabase()

	route := gin.Default()

	route.Use(func(c *gin.Context) {
		c.Set("db", db)

		//CORS Setting
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
	})

	route.GET("/api/lihat_harga_udang", controller.LihatDataUdang)
	route.POST("/api/tambah_harga_udang", controller.TambahDataUdang)

	route.Run(":8080")
}
