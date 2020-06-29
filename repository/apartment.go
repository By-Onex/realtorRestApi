package repository

import (
	"github.com/By-Onex/realtorRestApi/models"
	"github.com/jinzhu/gorm"
)

//ApartmentRepository Репозиторий апартоментов
type ApartmentRepository struct {
	*BaseRepository
}

//NewApartmentRepository возвращает новый репозиторий недвижимости
func NewApartmentRepository(db *gorm.DB) *ApartmentRepository {
	return &ApartmentRepository{&BaseRepository{"недвижимость", db}}
}

func (repo *ApartmentRepository) Search(aparts *[]models.Apartment) error {
	rows, err := repo.Raw("SELECT * FROM недвижимость WHERE стоимость > 2000000").Rows()
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
