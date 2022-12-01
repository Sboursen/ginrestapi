package controllers

import (
	"log"
	"net/http"

	"github.com/Sboursen/ginrestapi/models"
	"github.com/gin-gonic/gin"
)

type NewUser struct {
	Name    string `json:"name" binding:"required"`
	Email   int    `json:"email" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Country string `json:"country" binding:"required"`
}

type UserUpdate struct {
	Name    string `json:"name"`
	Email   int    `json:"email"`
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
	c.JSON(http.StatusOK, users)

}
