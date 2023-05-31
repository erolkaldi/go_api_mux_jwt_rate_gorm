package service

import (
	"github.com/erolkaldi/agency/pkg/models"
	"github.com/erolkaldi/agency/pkg/repository"
)

type UserService struct {
	repository *repository.Repository
}

func InitializeUserService(repository *repository.Repository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) MigrateUser() error {
	return s.repository.MigrateUser()
}
func (service *UserService) GetUserById(id int) (*models.User, error) {
	return service.repository.GetUserById(id)
}
func (service *UserService) GetUserByEmail(email string) (*models.User, error) {
	return service.repository.GetUserByEmail(email)
}
func (s *UserService) SaveUser(user *models.User) (*models.User, error) {
	return s.repository.SaveUser(user)
}
