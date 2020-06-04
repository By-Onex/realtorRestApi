package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/By-Onex/realtorRestApi/models"
	"github.com/By-Onex/realtorRestApi/repository"
	"github.com/By-Onex/realtorRestApi/utils"
	"github.com/gorilla/mux"
)

var favoriteRepo = repository.GetFavoriteRepo()

func AddFavoriteController(w http.ResponseWriter, r *http.Request) {
	var message map[string]interface{}

	params := mux.Vars(r)
	apartID, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
		message = utils.Message(false, "Неправильно задан параметр")
		utils.Respond(w, message)
		return
	}

	res, err := favoriteRepo.FindFirst(4, apartID)
	if err != nil {
		fmt.Println(err)
		message = utils.Message(false, "err")
		utils.Respond(w, message)
		return
	}
	if res == true {
		message = utils.Message(false, "Такая недвижимость уже добавлена")
		utils.Respond(w, message)
		return
	}

	res, err = apartRepo.CheckApart(apartID)
	if err != nil {
		fmt.Println(err)
		message = utils.Message(false, "Ошибка на сервере")
		utils.Respond(w, message)
		return
	}
	if res == false {
		message = utils.Message(false, "Такой недвижимости не существует")
		utils.Respond(w, message)
		return
	}

	err = favoriteRepo.Create(4, apartID)
	if err != nil {
		fmt.Println(err)
		message = utils.Message(false, "Ошибка на сервере")
	} else {
		message = utils.Message(true, "")
	}

	utils.Respond(w, message)
}

func AllFavoriteController(w http.ResponseWriter, r *http.Request) {

	favorite := []models.Apartment{}
	err := favoriteRepo.All(4, &favorite)
	if err != nil {
		fmt.Println(err)
		message := utils.Message(false, "error")
		utils.Respond(w, message)
		return
	}

	message := utils.Message(true, "")
	message["data"] = favorite
	utils.Respond(w, message)
}

func RemoveFavoriteController(w http.ResponseWriter, r *http.Request) {
	var message map[string]interface{}

	params := mux.Vars(r)
	apartID, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Println(err)
		msg := utils.Message(false, "Неправильный параметр")
		utils.Respond(w, msg)
		return
	}

	res, err := favoriteRepo.Remove(4, apartID)
	if err != nil {
		fmt.Println(err)
		message = utils.Message(false, "Ошибка удаления")
	} else if res == true {
		message = utils.Message(true, "")
	} else {
		message = utils.Message(false, "Такой недвижимоcти не найдено")
	}
	utils.Respond(w, message)
}
