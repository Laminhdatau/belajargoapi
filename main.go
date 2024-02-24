package main

import (
	"github.com/Laminhdatau/belajargoapi/controllers/productController"
	"github.com/Laminhdatau/belajargoapi/controllers/userController"
	"github.com/Laminhdatau/belajargoapi/config"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDb()
	r.GET("/api/v1/products", productController.Index)
	r.GET("/api/v1/products/:id", productController.Show)
	r.POST("/api/v1/products", productController.Create)
	r.PUT("/api/v1/products/:id", productController.Update)
	r.DELETE("/api/v1/products", productController.Delete)

	r.GET("/api/v1/users", userController.Index)
	r.GET("/api/v1/users/:id", userController.Show)
	r.POST("/api/v1/users", userController.Create)
	r.PUT("/api/v1/users/:id", userController.Update)
	r.DELETE("/api/v1/users", userController.Delete)

	r.Run()
}
