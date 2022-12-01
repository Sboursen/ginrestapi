package main

import (
	"fmt"
	"log"

	"github.com/Sboursen/ginrestapi/models"
)

func main() {
	fmt.Println("Hello")

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}
	db.DB()
}
