package models

//Request заявка на поиск недвижимости
type Request struct {
	ID int64 `json:"id"`

	UserID   int64 `json:"user_id" gorm:"column:сотрудник_id"`
	ClientID int64 `json:"client_id" gorm:"column:клиент_id"`

	MinPrice int `json:"min_price" gorm:"column:мин_цена"`
	MaxPrice int `json:"max_price" gorm:"column:макс_цена"`

	MinArea float32 `json:"min_area" gorm:"column:мин_площадь"`
	MaxArea float32 `json:"max_area" gorm:"column:макс_цена"`

	MinFloor int `json:"min_floor" gorm:"column:мин_этаж"`
	MaxFloor int `json:"max_floor" gorm:"column:макс_этаж"`

	MinStoreys int `json:"min_storeys" gorm:"column:мин_этажность"`
	MaxStoreys int `json:"max_storeys" gorm:"column:мах_этажность"`

	MinRoomCount int `json:"min_room_count" gorm:"column:мин_комнат"`
	MaxRoomCount int `json:"max_room_count" gorm:"column:макс_комнат"`

	District string `json:"district" gorm:"column:район"`
	Street   string `json:"street" gorm:"column:улица"`
	Num      string `json:"num" gorm:"column:номер"`
}
