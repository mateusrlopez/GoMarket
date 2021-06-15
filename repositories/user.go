package repositories

import (
	"github.com/mateusrlopez/go-market/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) RetrieveByEmail(email string, user *models.User) error {
	return r.DB.Where("email = ?", email).Take(user).Error
}
