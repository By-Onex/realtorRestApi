package repository

import (
	"github.com/By-Onex/realtorRestApi/models"
	"github.com/jinzhu/gorm"
)

//RequestRepository сотрудник
type RequestRepository struct {
	*gorm.DB
}

//NewRequestRepository возвращает новый репозиторий заявок
func NewRequestRepository(db *gorm.DB) *RequestRepository {
	return &RequestRepository{db}
}

//FindByID ищет заявку по id
func (repo *RequestRepository) FindByID(id int, req *models.Request) error {
	return repo.Raw("SELECT * FROM зявки WHERE id = ? LIMIT 1;", id).Scan(req).Error
}

//Create создает заявку
func (repo *RequestRepository) Create(newReq *models.Request) error {
	return repo.Table("зявки").Create(newReq).Error
}

//All возвращает список заявок
func (repo *RequestRepository) All(reqs *[]models.Request) error {
	rows, err := repo.Raw("SELECT * FROM заявки").Rows()
	defer rows.Close()

	if err != nil {
		return err
	}
	var req models.Request
	for rows.Next() {
		models.GetDB().ScanRows(rows, &req)
		*reqs = append(*reqs, req)
	}
	return nil
}

//Update обновляет данные заявки
func (repo *RequestRepository) Update(req *models.Request) error {
	return repo.Table("заявки").Update(req).Error
}
