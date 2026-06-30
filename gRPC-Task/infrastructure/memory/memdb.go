package memory

import "github.com/hashicorp/go-memdb"

type MemoryStore struct {
	db *memdb.MemDB
}

func CreateMemoryStore() (*MemoryStore, error) {
	db, err := memdb.NewMemDB(Schema)
	if err != nil {
		panic(err)
	}

	return &MemoryStore{db: db}, nil
}

func (s *MemoryStore) GetStore() *memdb.MemDB {
	return s.db
}
