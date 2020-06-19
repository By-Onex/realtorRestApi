package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/By-Onex/realtorRestApi/controllers"
	"github.com/By-Onex/realtorRestApi/models"
	"github.com/By-Onex/realtorRestApi/parser"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("start server")

	t, _ := parser.GetApartments(10)
	fmt.Println(t.Len())

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	models.ConnectDB()

	controllers.Initialize(models.GetDB())

	router := mux.NewRouter()
	router.Use(controllers.MiddlewareJwt)

	//Пользователь
	//Регистрация
	router.HandleFunc("/api/user/create", controllers.UserRegister).Methods(http.MethodPost)
	//Логин
	router.HandleFunc("/api/user/login", controllers.UserLogin).Methods(http.MethodPost)
	//Информация
	router.HandleFunc("/api/user/{id}", controllers.GetUserController).Methods(http.MethodGet)

	//Недвижимость
	//Весь список недвижимости
	router.HandleFunc("/api/apart/all", controllers.AllApartmentController).Methods(http.MethodGet)
	//Инфо про 1 недвижимость
	router.HandleFunc("/api/apart/{id}", controllers.GetApartmentController).Methods(http.MethodGet)

	//Избранная недвижимоть
	//Весь список недвижимости
	router.HandleFunc("/api/favorite/all", controllers.AllFavoriteController).Methods(http.MethodGet)
	//Добавление
	router.HandleFunc("/api/favorite/add/{id}", controllers.AddFavoriteController).Methods(http.MethodPost)
	//Удаление
	router.HandleFunc("/api/favorite/remove/{id}", controllers.RemoveFavoriteController).Methods(http.MethodDelete)

	port := os.Getenv("port")

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
