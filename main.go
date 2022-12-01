package main

import (
	"log"

	"github.com/Sboursen/ginrestapi/controllers"
	"github.com/Sboursen/ginrestapi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}
	db.DB()

	router := gin.Default()
	router.GET("/", controllers.GetUsers)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUserByID)
	router.POST("/users", controllers.PostUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	router.Run()
}
