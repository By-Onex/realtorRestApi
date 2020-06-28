package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/By-Onex/realtorRestApi/utils"
	"github.com/gorilla/mux"

	"github.com/By-Onex/realtorRestApi/models"
)

//GetApartmentController возвращает информацию о недвижимости
func GetApartmentController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		msg := utils.Message(false, "неправильный параметр")
		utils.Respond(w, msg)
		return
	}

	apart := &models.Apartment{}
	err = apartRepo.Get(id, apart)

	if err != nil {
		if err.Error() == "record not found" {
			msg := utils.Message(false, "не найдено")
			utils.Respond(w, msg)
			return
		}
	}

	msg := utils.Message(true, "")
	msg["data"] = apart
	utils.Respond(w, msg)
}

//SearchApartment поиск недвижимости в бд
func SearchApartment(w http.ResponseWriter, r *http.Request) {
	aparts := []models.Apartment{}
	err := apartRepo.Search(&aparts)
	if err != nil {
		fmt.Println(err)
	}

	msg := utils.Message(true, "")
	msg["data"] = aparts
	utils.Respond(w, msg)
}

//AllApartmentController возвращает информацию о всей недвижимости
func AllApartmentController(w http.ResponseWriter, r *http.Request) {
	aparts := []models.Apartment{}
	err := apartRepo.All(&aparts)
	if err != nil {
		fmt.Println(err)
	}

	msg := utils.Message(true, "")
	msg["data"] = aparts
	utils.Respond(w, msg)
}
