package repository

import (
	"github.com/By-Onex/realtorRestApi/models"
	"github.com/jinzhu/gorm"
)

//UserRepository сотрудник
type UserRepository struct {
	*BaseRepository
}

//NewUserRepository возвращает новый репозиторий сотрудника
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{&BaseRepository{"сотрудники", db}}
}

//UpdatePasswordUser обновление пароля сотрудника
func (repo *UserRepository) UpdatePasswordUser(id int64, password string) error {
	return repo.Exec("UPDATE сотрудники SET пароль = ? WHERE id = ? LIMIT 1;", password, id).Error
}

//FindUserByEmail находит пользователя по логину
func (repo *UserRepository) FindUserByEmail(email string, user *models.User) error {
	return repo.Raw("SELECT * FROM сотрудники WHERE email = ? LIMIT 1;", email).Scan(user).Error
}

/*
//FindUserByID находит сотрудника по идентификатору
func (repo *UserRepository) FindUserByID(id int, user *models.User) error {
	return repo.Raw("SELECT * FROM сотрудники WHERE id = ? LIMIT 1;", id).Scan(user).Error
}

//CreateUser создание нового сотрудника
func (repo *UserRepository) CreateUser(auth *models.Auth) error {
	return repo.Exec("INSERT INTO сотрудники (фио, email, телефон, пароль, роль_id) VALUES(?, ?, ?, ?, 2);",
		auth.FIO, auth.Email, auth.Number, auth.Password).Error
}

//UpdatePasswordUser обновление пароля сотрудника
func (repo *UserRepository) UpdatePasswordUser(id int, password string) error {
	return repo.Exec("UPDATE сотрудники SET пароль = ? WHERE id = ? LIMIT 1;", password, id).Error
}

//FindUserByEmail находит пользователя по логину
func (repo *UserRepository) FindUserByEmail(email string, user *models.User) error {
	return repo.Raw("SELECT * FROM сотрудники WHERE email = ? LIMIT 1;", email).Scan(user).Error
}

/*

//FindUserByUsername находит пользователя по логину
func (repo *UserRepository) FindUserByUsername(username string, user *models.User) error {
	return repo.Raw("SELECT * FROM usr WHERE username = ? LIMIT 1;", username).Scan(user).Error
}

//CreateUser создание нового пользователя
func (repo *UserRepository) CreateUser(username string, password string) error {
	return repo.Exec("INSERT INTO usr (username, password) VALUES(?, ?);", username, password).Error
}*/
