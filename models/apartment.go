package models

//Apartment Модель недвижимости
type Apartment struct {
	ID    int64   `json:"id"`
	Price int     `json:"price"`
	Area  float32 `json:"area"`

	Floor     int
	Storeys   int
	RoomCount int

	District string
	Street   string
	Num      string

	PageURL  string
	ImageURL string
}
