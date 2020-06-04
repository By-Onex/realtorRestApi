package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Message Создает ответ со статусом
func Message(status bool, message string) map[string]interface{} {
	if status == true {
		return map[string]interface{}{"status": status}
	}
	return map[string]interface{}{"status": status, "message": message}
}

//Respond Записывает сообщение json
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Println(err)
	}
}
