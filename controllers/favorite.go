package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/By-Onex/realtorRestApi/models"
	"github.com/By-Onex/realtorRestApi/utils"
	"github.com/gorilla/mux"
)

//AddFavoriteController добавеление указанной недвижимости в избранное
func AddFavoriteController(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(usr).(int64)

	var message map[string]interface{}

	params := mux.Vars(r)
	apartID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
		message = utils.Message(false, "Неправильно задан параметр")
		utils.Respond(w, message)
		return
	}

	res, err := favoriteRepo.FindFirst(userID, apartID)
	if err != nil {
		fmt.Println(err)
		message = utils.Message(false, "Ошибка на сервере")
		utils.Respond(w, message)
		return
	}
	if res == true {
		message = utils.Message(false, "Такая недвижимость уже добавлена")
		utils.Respond(w, message)
		return
	}

	err = apartRepo.FindByID(apartID, &models.Apartment{})
	if err != nil {
		if err.Error() == "record not found" {
			msg := utils.Message(false, fmt.Sprintf("Недвижимости с идентификатором %d не найдено", apartID))
			utils.Respond(w, msg)
			return
		}
		fmt.Println(err)
		message = utils.Message(false, "Ошибка на сервере")
		utils.Respond(w, message)
		return
	}

	err = favoriteRepo.Create(userID, apartID)
	if err != nil {
		fmt.Println(err)
		message = utils.Message(false, "Ошибка на сервере")
	} else {
		message = utils.Message(true, "")
	}

	utils.Respond(w, message)
}

//AllFavoriteController возвращает список избранной недвижимости
func AllFavoriteController(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(usr).(int64)

	favorite := []models.Apartment{}
	err := favoriteRepo.All(userID, &favorite)
	if err != nil {
		fmt.Println(err)
		message := utils.Message(false, "Ошибка на сервере")
		utils.Respond(w, message)
		return
	}

	message := utils.Message(true, "")
	message["data"] = favorite
	utils.Respond(w, message)
}

//RemoveFavoriteController удаление указанной недвижимости из избранного
func RemoveFavoriteController(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(usr).(int64)

	var message map[string]interface{}

	params := mux.Vars(r)
	apartID, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
		msg := utils.Message(false, "Неправильный параметр")
		utils.Respond(w, msg)
		return
	}

	res, err := favoriteRepo.Remove(userID, apartID)
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
