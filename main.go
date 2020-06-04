package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/By-Onex/realtorRestApi/controllers"
	"github.com/By-Onex/realtorRestApi/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("start server")

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	models.ConnectDB()

	controllers.Initialize(models.GetDB())

	router := mux.NewRouter()
	//Пользователь
	router.HandleFunc("/api/user/{id}", controllers.GetUserController).Methods("GET")
	//Недвижимость
	router.HandleFunc("/api/apart/all", controllers.AllApartmentController).Methods("GET")
	router.HandleFunc("/api/apart/{id}", controllers.GetApartmentController).Methods("GET")
	//Избранная недвижимоть
	router.HandleFunc("/api/favorite/all", controllers.AllFavoriteController).Methods("GET")
	router.HandleFunc("/api/favorite/add/{id}", controllers.AddFavoriteController).Methods("GET")
	router.HandleFunc("/api/favorite/remove/{id}", controllers.RemoveFavoriteController).Methods("GET")

	port := os.Getenv("port")

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
