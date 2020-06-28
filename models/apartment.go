package models

//Apartment модель недвижимости
type Apartment struct {
	ID    int64   `json:"id"`
	Price int     `json:"price" gorm:"column:стоимость"`
	Area  float64 `json:"area" gorm:"column:площадь"`

	Floor     int `json:"floor" gorm:"column:этаж"`
	Storeys   int `json:"storeys" gorm:"column:этажность"`
	RoomCount int `json:"room_count" gorm:"column:комнаты"`

	District string `json:"district" gorm:"column:район"`
	Street   string `json:"street" gorm:"column:улица"`
	Num      string `json:"num" gorm:"column:номер_дома"`

	PageURL  string `json:"page_url" gorm:"column:ссылка"`
	ImageURL string `json:"image_url" gorm:"column:изображение"`
}
