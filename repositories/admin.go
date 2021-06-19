package repositories

import (
	"github.com/mateusrlopez/go-market/models"
	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func (r *AdminRepository) RetrieveByEmail(email string, admin *models.Admin) error {
	return r.DB.Where("email = ?", email).Take(admin).Error
}
