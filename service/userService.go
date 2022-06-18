package service

import (
	"belajar-go-echo/model"
	"belajar-go-echo/repository"
)

type UserService interface {
	FindAll() ([]model.User, error)
	CreateUser(user model.User) (model.User, error)
}

type service struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]model.User, error) {
	return s.repository.FindAll()
}

func (s *service) CreateUser(user model.User) (model.User, error) {
	return s.repository.CreateUser(user)
}
