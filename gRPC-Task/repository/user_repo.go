package repository

import "server/infrastructure/models"

type UserRepository interface {
	CreateUser(u *models.User) error
	UpdateUser(u *models.User) error
	GetUser(userID string) (*models.User, error)
	DeleteUser(userID string) error
	GetAllUsers() ([]models.User, error)
}
