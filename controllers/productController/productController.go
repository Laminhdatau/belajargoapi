package productController

import (
	"encoding/json"
	"net/http"
	"github.com/Laminhdatau/belajargoapi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"astatus":  http.StatusOK,
		"products": products})
}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"astatus": http.StatusNotFound, "message": "Data tidak di temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"astatus": http.StatusInternalServerError, "message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"astatus": http.StatusOK, "product": product})
}
func Create(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"astatus": http.StatusBadRequest, "amessage": err.Error()})
		return
	}
	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"astatus": http.StatusOK, "product": product})
}
func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"astatus": http.StatusBadRequest, "message": err.Error()})
		return
	}
	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"astatus": http.StatusBadRequest, "message": "Tidak dapat di update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"astatus": http.StatusOK, "message": "Data di perbarui"})

}
func Delete(c *gin.Context) {
	var product models.Product
	var input struct {
		Id json.Number
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"astatus": http.StatusBadRequest, "message": err.Error()})
		return
	}
	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"astatus": http.StatusBadRequest, "message": "Tidak dapat di hapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"astatus": http.StatusOK, "message": "Data di hapus"})

}
