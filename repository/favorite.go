package repository

import (
	"github.com/By-Onex/realtorRestApi/models"
)

//FavoriteRepository Репозиторий избранной недвижимости
type FavoriteRepository struct{}

var favoriteRepo *FavoriteRepository

func init() {
	favoriteRepo = NewFavoriteRepo()
}

//NewFavoriteRepo Создает репозиторий
func NewFavoriteRepo() *FavoriteRepository {
	return &FavoriteRepository{}
}

//Create Создание
func (repo *FavoriteRepository) Create(userID int, apartID int) error {
	return models.GetDB().Exec("INSERT INTO favorite_apartments (user_id, apartment_id) VALUES(?, ?);", userID, apartID).Error
}

//Remove удаление записи
func (repo *FavoriteRepository) Remove(userID int, apartID int) (bool, error) {
	var count int
	err := models.GetDB().Raw(
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
func (repo *FavoriteRepository) FindFirst(userID int, apartID int) (bool, error) {
	var count int
	err := models.GetDB().Raw(
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
func (repo *FavoriteRepository) All(userID int, aparts *[]models.Apartment) error {
	rows, err := models.GetDB().Raw(
		`SELECT * FROM apartment a
		LEFT JOIN favorite_apartments f ON a.id = f.apartment_id
		WHERE f.user_id = ?`, userID).Rows()

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var apart models.Apartment
		err = models.GetDB().ScanRows(rows, &apart)
		if err != nil {
			return err
		}
		*aparts = append(*aparts, apart)
	}
	return nil
}

//GetFavoriteRepo Репозиторий
func GetFavoriteRepo() *FavoriteRepository {
	return favoriteRepo
}
