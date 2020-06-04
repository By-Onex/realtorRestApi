package repository

import (
	"github.com/By-Onex/realtorRestApi/models"
	"github.com/jinzhu/gorm"
)

//ApartmentRepository Репозиторий апартоментов
type ApartmentRepository struct {
	*gorm.DB
}

//NewApartmentRepository возвращает новый репозиторий недвижимости
func NewApartmentRepository(db *gorm.DB) *ApartmentRepository {
	return &ApartmentRepository{db}
}

//CheckApart проверяет существование недвижимости в БД
func (repo *ApartmentRepository) CheckApart(id int) (bool, error) {
	var count int
	err := repo.Raw("SELECT count(*) FROM apartment WHERE id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

//Get Найти недвижимоть по id
func (repo *ApartmentRepository) Get(id int, apart *models.Apartment) error {
	return repo.Raw("SELECT * FROM apartment WHERE id = ?", id).First(apart).Error
}

//All Найти недвижимоть
func (repo *ApartmentRepository) All(aparts *[]models.Apartment) error {
	rows, err := repo.Raw("SELECT * FROM apartment").Rows()
	defer rows.Close()

	if err != nil {
		return err
	}
	for rows.Next() {
		var apart models.Apartment
		models.GetDB().ScanRows(rows, &apart)
		*aparts = append(*aparts, apart)
	}
	return nil
}
