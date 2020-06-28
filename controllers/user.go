package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/By-Onex/realtorRestApi/models"
	"github.com/By-Onex/realtorRestApi/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

//GetUser ищет сотрудника по идентефикатору
func GetUser(w http.ResponseWriter, r *http.Request) {
	var message map[string]interface{}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		message = utils.Message(false, "Неправильный параметр")
		utils.Respond(w, message)
		return
	}

	user := &models.User{}
	err = userRepo.FindByID(id, user)

	if err != nil {
		if err.Error() == "record not found" {
			message = utils.Message(false, fmt.Sprintf("Пользователя с идентификатором %d не найдено", id))
			utils.Respond(w, message)
			return
		}
		fmt.Println(err)
		message = utils.Message(false, "Произошла ошибка на сервере")
		utils.Respond(w, message)
		return

	}

	message = utils.Message(true, "")
	message["data"] = user
	utils.Respond(w, message)
}

//UserLogin авторизация сотрудника
func UserLogin(w http.ResponseWriter, r *http.Request) {
	var message map[string]interface{}
	user := &models.User{}

	auth := &models.Auth{}

	err := json.NewDecoder(r.Body).Decode(auth) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		utils.Respond(w, utils.Message(false, "Неправильный запрос"))
		return
	}

	err = userRepo.FindUserByEmail(auth.Email, user)

	if err != nil {
		if err.Error() == "record not found" {
			message = utils.Message(false, "Пользователя с таким email не найдено")
			utils.Respond(w, message)
			return
		}
		fmt.Println(err)
		message = utils.Message(false, "Произошла ошибка на сервере")
		utils.Respond(w, message)
		return
	}

	//TODO добавить шифрование
	check := utils.CheckPassword(user.Password, auth.Password)

	if !check {
		message = utils.Message(false, "Неверный пароль")
		utils.Respond(w, message)
		return
	}

	//Создать токен JWT
	tk := &models.Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))

	message = utils.Message(true, "Успешно")
	message["token"] = tokenString
	utils.Respond(w, message)
}

//UserRegister регистрация сотрудника
func UserRegister(w http.ResponseWriter, r *http.Request) {
	var message map[string]interface{}
	user := &models.User{}

	auth := &models.Auth{}

	err := json.NewDecoder(r.Body).Decode(auth) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	err = userRepo.FindUserByEmail(auth.Email, user)
	if err != nil && err.Error() != "record not found" {
		fmt.Println(err)
		message = utils.Message(false, "Произошла ошибка на сервере")
		utils.Respond(w, message)
		return
	} else if user.Email != "" {
		message = utils.Message(false, "Пользователь с такой почтой уже зарегесрирован")
		utils.Respond(w, message)
		return
	}

	if auth.Password != auth.SecondPassword {
		message = utils.Message(false, "Пароли не совпадают")
		utils.Respond(w, message)
		return
	}

	if len(auth.Password) < 6 {
		message = utils.Message(false, "Короткий пароль")
		utils.Respond(w, message)
		return
	}

	err = userRepo.Create(auth)
	if err != nil {
		fmt.Println(err)
		message = utils.Message(false, "Произошла ошибка на сервере")
		utils.Respond(w, message)
		return
	}

	message = utils.Message(true, "Успешно")
	utils.Respond(w, message)
}

//UserUpdatePassword обновление пароля сотрудника
func UserUpdatePassword(w http.ResponseWriter, r *http.Request) {
	var message map[string]interface{}
	user := &models.User{}

	auth := &models.Auth{}

	err := json.NewDecoder(r.Body).Decode(auth) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	err = userRepo.FindUserByEmail(auth.Email, user)
	if err != nil && err.Error() != "record not found" {
		fmt.Println(err)
		message = utils.Message(false, "Произошла ошибка на сервере")
		utils.Respond(w, message)
		return
	} else if user.Email != "" {
		message = utils.Message(false, "Пользователь с таким логином уже существует")
		utils.Respond(w, message)
		return
	}

	if len(auth.Password) < 6 {
		message = utils.Message(false, "Короткий пароль")
		utils.Respond(w, message)
		return
	}

	err = userRepo.UpdatePasswordUser(user.ID, auth.Password)
	if err != nil {
		fmt.Println(err)
		message = utils.Message(false, "Произошла ошибка на сервере")
		utils.Respond(w, message)
		return
	}

	message = utils.Message(true, "Успешно")
	utils.Respond(w, message)
}
