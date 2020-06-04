package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/By-Onex/realtorRestApi/repository"
	"github.com/By-Onex/realtorRestApi/utils"
	"github.com/gorilla/mux"

	"github.com/By-Onex/realtorRestApi/models"
)

var apartRepo = repository.GetApartRepo()

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
