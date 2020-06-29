package repository

import (
	"github.com/By-Onex/realtorRestApi/models"
	"github.com/jinzhu/gorm"
)

//FavoriteRepository Репозиторий избранной недвижимости
type FavoriteRepository struct {
	*gorm.DB
}

//NewFavoriteRepository возвращает новый репозиторий избранной недвижимости
func NewFavoriteRepository(db *gorm.DB) *FavoriteRepository {
	return &FavoriteRepository{db}
}

//Create Создание
func (repo *FavoriteRepository) Create(userID int64, apartID int64) error {
	return repo.Exec("INSERT INTO избранная_недвижимость (сотрудник_id, недвижимость_id) VALUES(?, ?);", userID, apartID).Error
}

//Remove удаление записи
func (repo *FavoriteRepository) Remove(userID int64, apartID int64) (bool, error) {
	var count int
	err := repo.Raw(
		"WITH deleted AS (DELETE FROM избранная_недвижимость  WHERE сотрудник_id = ? AND недвижимость_id = ? RETURNING *) SELECT count(*) FROM deleted;",
		userID, apartID).Count(&count).Error

	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

//FindFirst проверить наличие такой же записи
func (repo *FavoriteRepository) FindFirst(userID int64, apartID int64) (bool, error) {
	var count int
	err := repo.Raw(
		"SELECT count(*) FROM избранная_недвижимость f WHERE f.сотрудник_id = ? AND f.недвижимость_id = ? LIMIT 1;", userID, apartID).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

//All Найти избранную недвижимоть пользователя
func (repo *FavoriteRepository) All(userID int64, aparts *[]models.Apartment) error {
	return repo.Raw(
		`SELECT * FROM недвижимость a
	LEFT JOIN избранная_недвижимость f ON a.id = f.недвижимость_id
	WHERE f.сотрудник_id = ?`, userID).Scan(aparts).Error

}
