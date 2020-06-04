package repository

import (
	"github.com/By-Onex/realtorRestApi/models"
	"github.com/jinzhu/gorm"
)

//UserRepository пользователь
type UserRepository struct {
	*gorm.DB
}

//NewUserRepository возвращает новый репозиторий пользователя
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

//FindUserByID находит пользователя по идентификатору
func (repo *UserRepository) FindUserByID(id int, user *models.User) error {
	return repo.Raw("SELECT * FROM usr WHERE id = ? LIMIT 1;", id).Scan(user).Error
}
