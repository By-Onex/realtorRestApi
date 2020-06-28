package controllers

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/By-Onex/realtorRestApi/models"
	"github.com/By-Onex/realtorRestApi/utils"
	"github.com/dgrijalva/jwt-go"
)

//MiddlewareJwt проверка токена и авторизация пользователя
func MiddlewareJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/api/user/create", "/api/user/login"}
		requestPath := r.URL.Path
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		var message map[string]interface{}
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" { //Токен отсутствует, возвращаем  403 http-код Unauthorized
			message = utils.Message(false, "Отсутствует токен авторизации")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, message)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //формат `Bearer {token-body}`
		if len(splitted) != 2 {
			message = utils.Message(false, "Неверный/неправильно сформированный токен авторизации")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, message)
			return
		}
		tokenPart := splitted[1] //Получаем вторую часть токена
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { //Неправильный токен, как правило, возвращает 403 http-код
			message = utils.Message(false, "Неверно сформированный токен аутентификации")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, message)
			return
		}

		if !token.Valid { //токен недействителен
			message = utils.Message(false, "Токен недействителен")
			w.WriteHeader(http.StatusForbidden)
			utils.Respond(w, message)
			return
		}

		//продолжаем выполнение запроса
		//fmt.Printf("User %d\n", tk.UserID)
		ctx := context.WithValue(r.Context(), usr, tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //передать управление следующему обработчику
	})

}
