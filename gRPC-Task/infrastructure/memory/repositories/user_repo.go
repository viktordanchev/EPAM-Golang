package repositories

import (
	"fmt"
	"server/infrastructure/memory"
	"server/infrastructure/memory/models"
)

type UserRepository struct {
	store *memory.MemoryStore
}

func NewUserRepository(store *memory.MemoryStore) *UserRepository {
	return &UserRepository{
		store: store,
	}
}

func (r *UserRepository) CreateUser(u models.User) {
	txn := r.store.GetStore().Txn(true)
	defer txn.Abort()

	if err := txn.Insert("user", u); err != nil {
		panic(err)
	}

	txn.Commit()
}

func (r *UserRepository) UpdateUser(u models.User) error {
	txn := r.store.GetStore().Txn(true)
	defer txn.Abort()

	existing, err := txn.First("user", "id", u.UserId)
	if err != nil {
		panic(err)
	}

	if existing == nil {
		return fmt.Errorf("User not found")
	}

	if err := txn.Insert("user", u); err != nil {
		panic(err)
	}

	txn.Commit()
	return nil
}

func (r *UserRepository) GetUser(userID string) (*models.User, error) {
	txn := r.store.GetStore().Txn(false)
	defer txn.Abort()

	raw, err := txn.First("user", "id", userID)
	if err != nil {
		panic(err)
	}

	if raw == nil {
		return nil, fmt.Errorf("User not found")
	}

	user := raw.(*models.User)

	return user, nil
}

func (r *UserRepository) DeleteUser(userID string) error {
	txn := r.store.GetStore().Txn(false)
	defer txn.Abort()

	user, err := txn.First("user", "id", userID)
	if err != nil {
		panic(err)
	}

	if user == nil {
		return fmt.Errorf("User not found")
	}

	txn.Delete("user", user)
	return nil
}

func (r *UserRepository) GetAllUsers() ([]*models.User, error) {
	txn := r.store.GetStore().Txn(false)
	defer txn.Abort()

	it, err := txn.Get("user", "id")
	if err != nil {
		panic(err)
	}

	var users []*models.User

	for obj := it.Next(); obj != nil; obj = it.Next() {
		users = append(users, obj.(*models.User))
	}

	return users, nil
}
