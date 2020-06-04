package models

//Request заявка на поиск недвижимости
type Request struct {
	ID int64 `json:"id"`

	MinPrice int `json:"min_price"`
	MaxPrice int `json:"max_price"`

	MinArea float32 `json:"min_area"`
	MaxArea float32 `json:"max_area"`

	MinFloor int `json:"min_floor"`
	MaxFloor int `json:"max_floor"`

	MinStoreys int `json:"min_storeys"`
	MaxStoreys int `json:"max_storeys"`

	MinRoomCount int `json:"min_room_count"`
	MaxRoomCount int `json:"max_room_count"`

	District string `json:"district"`
	Street   string `json:"street"`
	Num      string `json:"num"`
}
