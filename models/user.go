package models

//User пользователь
type User struct {
	ID int64 `json:"id"`

	FIO   string `json:"fio"`
	Email string `json:"email"`
}
