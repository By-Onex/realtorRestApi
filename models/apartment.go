package models

//Apartment модель недвижимости
type Apartment struct {
	ID    int64   `json:"id"`
	Price int     `json:"price"`
	Area  float32 `json:"area"`

	Floor     int `json:"floor"`
	Storeys   int `json:"storeys"`
	RoomCount int `json:"room_count"`

	District string `json:"district"`
	Street   string `json:"street"`
	Num      string `json:"num"`

	PageURL  string `json:"page_url"`
	ImageURL string `json:"image_url"`
}
