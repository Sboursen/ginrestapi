package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() (*gorm.DB, error) {

	var envs map[string]string
	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatal("Error loading database credentials from .env file")
	}

	hostname := envs["HOST"]
	username := envs["USER"]
	password := envs["PASSWORD"]
	dbname := envs["DBNAME"]
	port := envs["PORT"]

	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s", hostname, username, password, dbname, port)

	sqlDB, _ := sql.Open("pgx", dsn)
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.AutoMigrate(&User{}); err == nil && db.Migrator().HasTable(&User{}) {
		var rowCount int64
		db.Table("users").Count(&rowCount)

		if rowCount == 0 {
			data, err := os.ReadFile("./models/people.json")
			if err == nil {
				var users []User
				err := json.Unmarshal(data, &users)

				if err == nil {
					db.Model(&User{}).Create(users)
				}
			}
		}
	}

	return db, err
}
