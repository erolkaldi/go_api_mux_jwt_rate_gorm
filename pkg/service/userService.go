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
func (s *UserService) RegisterUser(register *models.Register) *models.Response {
	user := models.User{Name: register.Name, Email: register.Email, Password: register.Password}
	returnUser, err := s.repository.SaveUser(&user)
	if err != nil {
		return &models.Response{IsSuccessfull: false, Message: err.Error()}
	}
	email := CreateRegisterEmail(returnUser)

	_, err = CreateOutBox(JsonToString(email), s.repository)
	if err != nil {
		return &models.Response{IsSuccessfull: false, Message: err.Error()}
	}

	return &models.Response{IsSuccessfull: true, Message: "OK"}
}
