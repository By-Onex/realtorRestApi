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
func (repo *FavoriteRepository) Create(userID int64, apartID int) error {
	return repo.Exec("INSERT INTO favorite_apartments (user_id, apartment_id) VALUES(?, ?);", userID, apartID).Error
}

//Remove удаление записи
func (repo *FavoriteRepository) Remove(userID int64, apartID int) (bool, error) {
	var count int
	err := repo.Raw(
		"WITH deleted AS (DELETE FROM favorite_apartments  WHERE user_id = ? AND apartment_id = ? RETURNING *) SELECT count(*) FROM deleted;",
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
func (repo *FavoriteRepository) FindFirst(userID int64, apartID int) (bool, error) {
	var count int
	err := repo.Raw(
		"SELECT count(*) FROM favorite_apartments f WHERE f.user_id = ? AND f.apartment_id = ? LIMIT 1;", userID, apartID).Count(&count).Error
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
	rows, err := repo.Raw(
		`SELECT * FROM apartment a
		LEFT JOIN favorite_apartments f ON a.id = f.apartment_id
		WHERE f.user_id = ?`, userID).Rows()

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var apart models.Apartment
		err = repo.ScanRows(rows, &apart)
		if err != nil {
			return err
		}
		*aparts = append(*aparts, apart)
	}
	return nil
}
