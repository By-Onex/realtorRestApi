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

//GetClient возвращает инфрмацию о клиенте
func GetClient(w http.ResponseWriter, r *http.Request) {
	var message map[string]interface{}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		message = utils.Message(false, "Неправильный параметр")
		utils.Respond(w, message)
		return
	}

	client := &models.Client{}
	err = clientRepo.FindByID(int64(id), client)

	if err != nil {
		if err.Error() == "record not found" {
			message = utils.Message(false, fmt.Sprintf("Клиента с идентификатором %d не найдено", id))
			utils.Respond(w, message)
			return
		}
		fmt.Println(err)
		message = utils.Message(false, "Произошла ошибка на сервере")
		utils.Respond(w, message)
		return
	}

	message = utils.Message(true, "")
	message["data"] = client
	utils.Respond(w, message)
}

//AllClients возвращает список клиентов
func AllClients(w http.ResponseWriter, r *http.Request) {
	var message map[string]interface{}

	clients := []models.Client{}
	err := clientRepo.All(&clients)
	if err != nil {
		fmt.Println(err)
		message = utils.Message(false, "Произошла ошибка на сервере")
		utils.Respond(w, message)
		return
	}

	msg := utils.Message(true, "")
	msg["data"] = clients
	utils.Respond(w, msg)
}

//CreateClient создает клиента
func CreateClient(w http.ResponseWriter, r *http.Request) {
	var message map[string]interface{}
	client := &models.Client{}

	err := json.NewDecoder(r.Body).Decode(client)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Некорректные данные"))
		return
	}

	if client.Email == "" && client.Number == "" {
		message = utils.Message(false, "Хотя бы одно из полей телефон или email должно быть заполнено")
		utils.Respond(w, message)
		return
	}

	err = clientRepo.Create(client)
	if err != nil {
		fmt.Println(err)
		message = utils.Message(false, "Произошла ошибка на сервере")
		utils.Respond(w, message)
		return
	}

	message = utils.Message(true, "Успешно")
	utils.Respond(w, message)
}
