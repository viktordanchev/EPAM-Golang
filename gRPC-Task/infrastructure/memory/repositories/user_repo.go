package repositories

import (
	"fmt"
	"server/infrastructure/memory/models"

	"github.com/hashicorp/go-memdb"
)

type UserRepository struct {
	store *memdb.MemDB
}

func NewUserRepository(store *memdb.MemDB) *UserRepository {
	return &UserRepository{
		store: store,
	}
}

func (r *UserRepository) CreateUser(u models.User) error {
	txn := r.store.Txn(true)
	defer txn.Abort()

	if err := txn.Insert("user", u); err != nil {
		return fmt.Errorf("insert user failed: %w", err)
	}

	txn.Commit()

	return nil
}

func (r *UserRepository) UpdateUser(u models.User) error {
	txn := r.store.Txn(true)
	defer txn.Abort()

	existing, err := txn.First("user", "id", u.UserId)
	if err != nil {
		return err
	}

	if existing == nil {
		return fmt.Errorf("User not found")
	}

	if err := txn.Insert("user", u); err != nil {
		return err
	}

	txn.Commit()
	return nil
}

func (r *UserRepository) GetUser(userID string) (models.User, error) {
	txn := r.store.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("user", "id", userID)
	if err != nil {
		return models.User{}, fmt.Errorf("Receive user failed: %w", err)
	}

	if raw == nil {
		return models.User{}, fmt.Errorf("User not found")
	}

	user := raw.(models.User)

	return user, nil
}

func (r *UserRepository) DeleteUser(userID string) error {
	txn := r.store.Txn(true)
	defer txn.Abort()

	user, err := txn.First("user", "id", userID)
	if err != nil {
		return err
	}

	if user == nil {
		return fmt.Errorf("user not found")
	}

	if err := txn.Delete("user", user); err != nil {
		return err
	}

	txn.Commit()
	return nil
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	txn := r.store.Txn(false)
	defer txn.Abort()

	it, err := txn.Get("user", "id")
	if err != nil {
		return nil, err
	}

	var users []models.User

	for obj := it.Next(); obj != nil; obj = it.Next() {
		users = append(users, obj.(models.User))
	}

	return users, nil
}
