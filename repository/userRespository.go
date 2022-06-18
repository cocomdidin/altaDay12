package repository

import (
	"belajar-go-echo/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	CreateUser(user model.User) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}