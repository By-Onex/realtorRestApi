package repository

import (
	"github.com/jinzhu/gorm"
)

//BaseRepository базовый репозиторий
type BaseRepository struct {
	table string
	*gorm.DB
}

//NewBaseRepository возвращает новый базовый репозиторий
func NewBaseRepository(table string, db *gorm.DB) *BaseRepository {
	return &BaseRepository{table, db}
}

//FindByID поиск сущности по id
func (repo *BaseRepository) FindByID(id int64, out interface{}) error {
	return repo.Raw("SELECT * FROM "+repo.table+" WHERE id = ? LIMIT 1;", id).Scan(out).Error
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
func (repo *BaseRepository) Update(entity interface{}) error {
	return repo.Table(repo.table).Update(entity).Error
}
