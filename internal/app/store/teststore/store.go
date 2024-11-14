package teststore

import (
	"database/sql"
	"github.com/MirrexOne/http-rest-api/internal/app/model"
	"github.com/MirrexOne/http-rest-api/internal/app/store"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}

	return s.userRepository
}
