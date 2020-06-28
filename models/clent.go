package models

//Client клиент
type Client struct {
	ID     int64  `json:"id"`
	FIO    string `json:"fio" gorm:"column:фио"`
	Email  string `json:"email"`
	Number string `json:"number" gorm:"column:телефон"`
}
