package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/By-Onex/realtorRestApi/models"
	"github.com/By-Onex/realtorRestApi/utils"
	"github.com/gorilla/mux"
)

//GetUserController ищет пользователя по идентефикатору
func GetUserController(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		msg := utils.Message(false, "Неправильный параметр")
		utils.Respond(w, msg)
		return
	}

	user := &models.User{}
	err = userRepo.FindUserByID(id, user)

	if err != nil {
		if err.Error() == "record not found" {
			msg := utils.Message(false, fmt.Sprintf("Пользователя с идентификатором %d не найдено", id))
			utils.Respond(w, msg)
			return
		}
	}

	msg := utils.Message(true, "")
	msg["data"] = user
	utils.Respond(w, msg)
}
