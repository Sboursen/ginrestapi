package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sboursen/ginrestapi/models"
	"github.com/gin-gonic/gin"
)

type NewUser struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Country string `json:"country" binding:"required"`
}

type UserUpdate struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Country string `json:"country"`
}

func GetUsers(c *gin.Context) {

	var users []models.User

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, users)

}

func GetUserByID(c *gin.Context) {
	var user models.User

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}
	if err := db.Where("id= ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User with id %s not found", c.Param("id"))})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func PostUser(c *gin.Context) {

	var newUser NewUser

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: newUser.Name, Email: newUser.Email, Phone: newUser.Phone, Country: newUser.Country}

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {

	var user models.User

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User with id %s not found", c.Param("id"))})
		return
	}

	var newUser UserUpdate

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&user).Updates(models.User{Name: newUser.Name, Email: newUser.Email, Phone: newUser.Phone, Country: newUser.Country}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, user)

}

func DeleteUser(c *gin.Context) {

	var user models.User

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User with id %s not found", c.Param("id"))})
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted"})

}
