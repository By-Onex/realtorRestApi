package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	//github.com/jinzhu/gorm/dialects/postgres драйвер для postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

//ConnectDB подключение к базе данных
func ConnectDB() {
	fmt.Println("connect bd")

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURL := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	var err error
	db, err = gorm.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
}

//GetDB возвращает подключение к базе данных
func GetDB() *gorm.DB {
	return db
}
