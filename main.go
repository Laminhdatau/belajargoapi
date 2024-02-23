package main

import (
	"github.com/Laminhdatau/belajargoapi/controllers/productController"
	"github.com/Laminhdatau/belajargoapi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDb()
	r.GET("/api/v1/products", productController.Index)
	r.GET("/api/v1/products/:id", productController.Show)
	r.POST("/api/v1/products", productController.Create)
	r.PUT("/api/v1/products/:id", productController.Update)
	r.DELETE("/api/v1/products", productController.Delete)

	r.Run()
}
