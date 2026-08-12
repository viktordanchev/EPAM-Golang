package database

import (
	"server/infrastructure/database/models"

	"gorm.io/gorm"
)

func MigrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Project{},
		&models.Issue{},
	)
}
