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

/*
CREATE TABLE клиенты (
	id serial PRIMARY KEY,
	фио TEXT NOT NULL,
	телефон TEXT,
	email TEXT NOT NULL);

CREATE TABLE сотрудники (
	id serial PRIMARY KEY,
	фио TEXT NOT NULL,
	телефон TEXT,
	email TEXT,
	пароль TEXT);

CREATE TABLE заявки (
	id serial PRIMARY KEY,
	клиент_id INTEGER NOT NULL REFERENCES клиенты (id),
	сотрудник_id INTEGER NOT NULL REFERENCES сотрудники (id),

	мин_стоимость INTEGER NOT NULL,
	макс_стоимость INTEGER NOT NULL,

	мин_площадь REAL NOT NULL,
	макс_площадь REAL NOT NULL,

	мин_этаж INTEGER NOT NULL,
	макс_этаж INTEGER NOT NULL,

	мин_этажность INTEGER NOT NULL,
	макс_этажность INTEGER NOT NULL,

	мин_комнат INTEGER NOT NULL,
	макс_комнат INTEGER NOT NULL,

	район TEXT NOT NULL,
	улица TEXT NOT NULL,
	номер_дома TEXT NOT NULL);

CREATE TABLE недвижимость (
	id serial PRIMARY KEY,
	стоимость INTEGER NOT NULL,
	площадь REAL NOT NULL,
	этаж INTEGER NOT NULL,
	этажность INTEGER NOT NULL,
	комнаты INTEGER NOT NULL,
	район TEXT NOT NULL,
	улица TEXT NOT NULL,
	номер_дома TEXT NOT NULL);

CREATE TABLE найденная_недвижимость (
	id serial PRIMARY KEY,
	заявка_id INTEGER NOT NULL REFERENCES заявки (id),
	недвижиость_id INTEGER NOT NULL REFERENCES недвижимость (id));
*/

//GetDB возвращает подключение к базе данных
func GetDB() *gorm.DB {
	return db
}
