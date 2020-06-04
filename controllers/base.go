package controllers

import (
	"github.com/By-Onex/realtorRestApi/repository"
	"github.com/jinzhu/gorm"
)

var (
	userRepo     *repository.UserRepository
	apartRepo    *repository.ApartmentRepository
	favoriteRepo *repository.FavoriteRepository
)

//Initialize создание репозиториев
func Initialize(db *gorm.DB) {
	userRepo = repository.NewUserRepository(db)
	apartRepo = repository.NewApartmentRepository(db)
	favoriteRepo = repository.NewFavoriteRepository(db)
}