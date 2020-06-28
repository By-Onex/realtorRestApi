package repository

import (
	"github.com/By-Onex/realtorRestApi/models"
	"github.com/jinzhu/gorm"
)

//ClientRepository репозиторий клиентов
type ClientRepository struct {
	*gorm.DB
}

//NewClientRepository возвращает новый репозиторий клиентов
func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{db}
}

//FindByID ищет клиента по id
func (repo *ClientRepository) FindByID(id int, client *models.Client) error {
	return repo.Raw("SELECT * FROM клиенты WHERE id = ? LIMIT 1;", id).Scan(client).Error
}

//Create создает клиента
func (repo *ClientRepository) Create(newClient *models.Client) error {
	return repo.Table("клиенты").Create(newClient).Error
}

//All возвращает список клиентов
func (repo *ClientRepository) All(clients *[]models.Client) error {
	rows, err := repo.Raw("SELECT * FROM клиенты").Rows()
	defer rows.Close()

	if err != nil {
		return err
	}
	var client models.Client
	for rows.Next() {
		models.GetDB().ScanRows(rows, &client)
		*clients = append(*clients, client)
	}
	return nil
}

//Update обновляет данные клиента
func (repo *ClientRepository) Update(client *models.Client) error {
	return repo.Table("клиенты").Update(client).Error
}
