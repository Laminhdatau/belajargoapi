package userController

import (
	"encoding/json"
	"net/http"
	"github.com/Laminhdatau/belajargoapi/config"
	"github.com/Laminhdatau/belajargoapi/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var users []models.User

	config.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"astatus":  http.StatusOK,
		"users": users})
}

func Show(c *gin.Context) {
	var User models.User
	id := c.Param("id")
	if err := config.DB.First(&User, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"astatus": http.StatusNotFound, "message": "Data tidak di temukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"astatus": http.StatusInternalServerError, "message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"astatus": http.StatusOK, "User": User})
}
func Create(c *gin.Context) {
	var User models.User
	if err := c.ShouldBindJSON(&User); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"astatus": http.StatusBadRequest, "amessage": err.Error()})
		return
	}
	config.DB.Create(&User)
	c.JSON(http.StatusCreated, gin.H{"astatus": http.StatusCreated, "User": User})
}
func Update(c *gin.Context) {
	var User models.User
	id := c.Param("id")
	if err := c.ShouldBindJSON(&User); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"astatus": http.StatusBadRequest, "message": err.Error()})
		return
	}
	if config.DB.Model(&User).Where("id = ?", id).Updates(&User).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"astatus": http.StatusBadRequest, "message": "Tidak dapat di update"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"astatus": http.StatusCreated, "message": "Data di perbarui"})

}
func Delete(c *gin.Context) {
	var User models.User
	var input struct {
		Id json.Number
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"astatus": http.StatusBadRequest, "message": err.Error()})
		return
	}
	id, _ := input.Id.Int64()
	if config.DB.Delete(&User, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"astatus": http.StatusBadRequest, "message": "Tidak dapat di hapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"astatus": http.StatusOK, "message": "Data di hapus"})

}
