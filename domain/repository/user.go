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
	SaveUser(user model.Users) (model.Users, error)
	Update(user model.Users) (model.Users, error)
	CheckEmail(email string) (bool, error)
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

func (r *userRepository) SaveUser(user model.Users) (model.Users, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Update(user model.Users) (model.Users, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) CheckEmail(email string) (bool, error) {
	err := r.db.Where("email = ?", email).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
