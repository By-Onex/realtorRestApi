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

//FindUserByUsername находит пользователя по логину
func (repo *UserRepository) FindUserByUsername(username string, user *models.User) error {
	return repo.Raw("SELECT * FROM usr WHERE username = ? LIMIT 1;", username).Scan(user).Error
}

//CreateUser создание нового пользователя
func (repo *UserRepository) CreateUser(username string, password string) error {
	return repo.Exec("INSERT INTO usr (username, password) VALUES(?, ?);", username, password).Error
}
