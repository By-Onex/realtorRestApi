package controllers

import (
	"net/http"

	"github.com/By-Onex/realtorRestApi/utils"
)

//IndexController тестовая
func IndexController(w http.ResponseWriter, r *http.Request) {
	message := utils.Message(true, "hello")
	utils.Respond(w, message)
}
