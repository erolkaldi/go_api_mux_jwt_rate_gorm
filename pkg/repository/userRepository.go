package repository

import (
	"github.com/erolkaldi/agency/pkg/models"
)

func (rp *Repository) Migrate() error {
	return rp.db.AutoMigrate(&models.User{}, &models.OutBox{})
}

func (rp *Repository) GetUserById(id int) (*models.User, error) {
	user := new(models.User)
	err := rp.db.Where(`id = ?`, id).First(&user).Error
	return user, err
}

func (rp *Repository) SaveUser(user *models.User) (*models.User, error) {
	err := rp.db.Save(&user).Error
	return user, err
}

func (rp *Repository) GetUserByEmail(email string) (*models.User, error) {
	user := new(models.User)
	err := rp.db.Where(`email = ?`, email).First(&user).Error
	return user, err
}
