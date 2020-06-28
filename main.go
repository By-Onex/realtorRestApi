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

	/*t, _ := parser.GetApartments(10)
	fmt.Println(t.Len())
	for e := t.Front(); e != nil; e = e.Next() {
		rep := repository.NewApartmentRepository(models.GetDB())
		err := rep.AddApart(e.Value.(*models.Apartment))
		if err != nil {
			fmt.Println(e.Value)
		}
	}
	router.Use(controllers.MiddlewareJwt)*/

	//Сотрудник
	//Регистрация
	router.HandleFunc("/api/user/create", controllers.UserRegister).Methods(http.MethodPost)
	//Логин
	router.HandleFunc("/api/user/login", controllers.UserLogin).Methods(http.MethodPost)
	//Обновление пароля
	router.HandleFunc("/api/user/{id}/password", controllers.UserUpdatePassword).Methods(http.MethodPut)
	//Информация
	router.HandleFunc("/api/user/{id}", controllers.GetUser).Methods(http.MethodGet)

	//Недвижимость
	//Весь список недвижимости
	router.HandleFunc("/api/apart/all", controllers.AllApartmentController).Methods(http.MethodGet)
	//Поиск недвижимости
	router.HandleFunc("/api/apart/search", controllers.SearchApartment).Methods(http.MethodGet)
	//Инфо про 1 недвижимость
	router.HandleFunc("/api/apart/{id}", controllers.GetApartmentController).Methods(http.MethodGet)

	//Избранная недвижимоть
	//Весь список недвижимости
	router.HandleFunc("/api/favorite/all", controllers.AllFavoriteController).Methods(http.MethodGet)
	//Добавление
	router.HandleFunc("/api/favorite/{id}", controllers.AddFavoriteController).Methods(http.MethodPost)
	//Удаление
	router.HandleFunc("/api/favorite/{id}", controllers.RemoveFavoriteController).Methods(http.MethodDelete)

	port := os.Getenv("port")

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
