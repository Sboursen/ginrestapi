package models

import (
	"database/sql"
	"fmt"
	"log"

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
	fmt.Println(dsn)
	sqlDB, _ := sql.Open("pgx", dsn)
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&User{}); err != nil {
		log.Println(err)
	}

	return db, err
}
