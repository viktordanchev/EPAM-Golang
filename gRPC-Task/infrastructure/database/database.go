package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {

	dsn := "host=postgres user=admin password=123456 dbname=grpc-server port=5432"

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
