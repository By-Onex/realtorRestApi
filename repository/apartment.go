package repository

import "github.com/By-Onex/realtorRestApi/models"

//ApartmentRepository Репозиторий апартоментов
type ApartmentRepository struct{}

var apartRepo *ApartmentRepository

func init() {
	apartRepo = NewApartmentRepo()
}

//NewApartmentRepo Конструктор
func NewApartmentRepo() *ApartmentRepository {
	return &ApartmentRepository{}
}

//Create Создание
func (repo *ApartmentRepository) Create(params ...interface{}) error {
	return models.GetDB().Exec("INSERT INTO apartments () VALUES('?');", params).Error
}

//CheckApart проверяет существование недвижимости в БД
func (repo *ApartmentRepository) CheckApart(id int) (bool, error) {
	var count int
	err := models.GetDB().Raw("SELECT count(*) FROM apartment WHERE id = ?", id).Count(&count).Error
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
	return models.GetDB().Raw("SELECT * FROM apartment WHERE id = ?", id).Scan(apart).Error
}

//All Найти недвижимоть
func (repo *ApartmentRepository) All(aparts *[]models.Apartment) error {
	rows, err := models.GetDB().Raw("SELECT * FROM apartment").Rows()
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

//GetApartRepo Репозиторий
func GetApartRepo() *ApartmentRepository {
	return apartRepo
}
