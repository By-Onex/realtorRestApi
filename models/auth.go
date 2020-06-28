package models

type Auth struct {
	FIO            string `json:"fio"`
	Number         string `json:"number"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	SecondPassword string `json:"second_password"`
}
