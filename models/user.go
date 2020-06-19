package models

//User сотрудник
type User struct {
	ID       int64  `json:"id"`
	FIO      string `json:"fio"`
	Email    string `json:"email"`
	Number   string `json:"number"`
	Password string `json:"-"`
}
