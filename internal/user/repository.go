package user

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *User) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (repo repository) CreateUser(user *User) error {
	result := repo.db.Create(user)
	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return errors.New("User with this email is already exists")
	}
	return nil
}
