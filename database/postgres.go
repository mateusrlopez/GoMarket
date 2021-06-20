package database

import (
	"fmt"

	"github.com/mateusrlopez/go-market/models"
	"github.com/mateusrlopez/go-market/settings"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(formatPostgresDSN()), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error opening connection with postgres database: %s", err)
		return nil
	}

	db.AutoMigrate(&models.User{})

	return db
}

func formatPostgresDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		settings.Settings.Database.Host,
		settings.Settings.Database.UserName,
		settings.Settings.Database.Password,
		settings.Settings.Database.Name,
		settings.Settings.Database.Port,
	)
}
