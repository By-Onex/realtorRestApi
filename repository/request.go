package repository

import (
	"github.com/jinzhu/gorm"
)

//RequestRepository сотрудник
type RequestRepository struct {
	*BaseRepository
}

//NewRequestRepository возвращает новый репозиторий заявок
func NewRequestRepository(db *gorm.DB) *RequestRepository {
	return &RequestRepository{&BaseRepository{"заявки", db}}
}
