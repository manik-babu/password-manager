package user

import (
	"errors"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) CreateUser(user *User) error {
	result := repo.db.Create(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return errors.New("User with this email is already exists")
		}
		return result.Error
	}
	return nil
}
func (repo *repository) GetUserByEmail(email string) (*User, error) {
	var user User
	result := repo.db.Where(&User{
		Email: email,
	}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
