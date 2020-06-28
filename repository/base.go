package repository

import (
	"github.com/By-Onex/realtorRestApi/models"
	"github.com/jinzhu/gorm"
)

//BaseRepository базовый репозиторий
type BaseRepository struct {
	table string
	*gorm.DB
}

//NewBaseRepository возвращает новый репозиторий клиентов
func NewBaseRepository(table string, db *gorm.DB) *BaseRepository {
	return &BaseRepository{table, db}
}

//FindByID ищет сущность по id
func (repo *BaseRepository) FindByID(id int, entity interface{}) error {
	return repo.Raw("SELECT * FROM "+repo.table+" WHERE id = ? LIMIT 1;", id).Scan(entity).Error
}

//Create создает сущность
func (repo *BaseRepository) Create(entity interface{}) error {
	return repo.Table(repo.table).Create(entity).Error
}

//All возвращает список сущностей
func (repo *BaseRepository) All(out interface{}) error {
	return repo.Table(repo.table).Find(out).Error
}

//Update обновляет данные сущности
func (repo *BaseRepository) Update(client *models.Client) error {
	return repo.Table(repo.table).Update(client).Error
}
