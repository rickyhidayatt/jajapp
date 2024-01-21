package repository

import (
	"github.com/jajapp/domain/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}
type UserRepositoryInterface interface {
	FindEmail(email string) (model.Users, error)
	FindByUuid(uuid string) (model.Users, error)
}

func NewUserRepository(g *gorm.DB) UserRepositoryInterface {
	return &userRepository{db: g}
}

func (r *userRepository) FindEmail(email string) (model.Users, error) {
	var user model.Users
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByUuid(uuid string) (model.Users, error) {
	var user model.Users
	err := r.db.Where("uuid = ?", uuid).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
