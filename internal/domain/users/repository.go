package users

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user *User) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

type repository struct {
	db *gorm.DB
}

var ErrorAlreadyExist = errors.New("user with this email already exist")

func (r *repository) CreateUser(user *User) error {
	result := r.db.Create(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return ErrorAlreadyExist
		}
	}
	return nil
}
