package repository

import (
	"fmt"

	"server/infrastructure/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(u *models.User) error {
	if err := r.db.Create(u).Error; err != nil {
		return fmt.Errorf("create user failed: %w", err)
	}

	return nil
}

func (r *UserRepository) UpdateUser(u *models.User) error {
	result := r.db.
		Model(&models.User{}).
		Where("user_id = ?", u.UserId).
		Updates(u)

	if result.Error != nil {
		return fmt.Errorf("update user failed: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func (r *UserRepository) GetUser(userID string) (*models.User, error) {
	var user models.User

	if err := r.db.
		Where("user_id = ?", userID).
		First(&user).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found")
		}

		return nil, fmt.Errorf("get user failed: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) DeleteUser(userID string) error {
	result := r.db.
		Where("user_id = ?", userID).
		Delete(&models.User{})

	if result.Error != nil {
		return fmt.Errorf("delete user failed: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User

	if err := r.db.
		Order("user_id").
		Find(&users).Error; err != nil {

		return nil, fmt.Errorf("get all users failed: %w", err)
	}

	return users, nil
}
