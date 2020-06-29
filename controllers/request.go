package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/By-Onex/realtorRestApi/models"
	"github.com/By-Onex/realtorRestApi/utils"
	"github.com/gorilla/mux"
)

//GetRequest возвращает инфрмацию о заявке
func GetRequest(w http.ResponseWriter, r *http.Request) {
	var message map[string]interface{}

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		message = utils.Message(false, "Неправильный параметр")
		utils.Respond(w, message)
		return
	}

	req := &models.Request{}
	err = requestRepo.FindByID(int64(id), req)

	if err != nil {
		if err.Error() == "record not found" {
			message = utils.Message(false, fmt.Sprintf("Заявки с идентификатором %d не найдено", id))
			utils.Respond(w, message)
			return
		}
		fmt.Println(err)
		message = utils.Message(false, "Произошла ошибка на сервере")
		utils.Respond(w, message)
		return
	}

	message = utils.Message(true, "")
	message["data"] = req
	utils.Respond(w, message)
}

//AllRequsets возвращает список заявок
func AllRequsets(w http.ResponseWriter, r *http.Request) {
	var message map[string]interface{}

	reqs := []models.Request{}
	err := requestRepo.All(&reqs)
	if err != nil {
		fmt.Println(err)
		message = utils.Message(false, "Произошла ошибка на сервере")
		utils.Respond(w, message)
		return
	}

	msg := utils.Message(true, "")
	msg["data"] = reqs
	utils.Respond(w, msg)
}

//CreateRequest создает заявку
func CreateRequest(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(usr).(int64)

	var message map[string]interface{}
	req := &models.Request{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Некорректные данные"))
		return
	}

	req.UserID = userID

	var client models.Client
	err = clientRepo.FindByID(req.ClientID, &client)
	if err != nil {
		if err.Error() != "record not found" {
			message = utils.Message(false, fmt.Sprintf("Клиента с идентификатором %d не найдено", req.ClientID))
			utils.Respond(w, message)
			return
		}
		message = utils.Message(false, "Произошла ошибка на сервере")
		utils.Respond(w, message)
		return
	}

	err = requestRepo.Create(req)
	if err != nil {
		fmt.Println(err)
		message = utils.Message(false, "Произошла ошибка на сервере")
		utils.Respond(w, message)
		return
	}

	message = utils.Message(true, "Успешно")
	utils.Respond(w, message)
}
