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
	query := `
	INSERT INTO users (created_at, name, email, password, address, phone_number, is_seller, is_driver, uuid, nik)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING created_at, name, email, password, address, phone_number, is_seller, is_driver, uuid, nik
    `
	rows, err := r.db.Raw(query,
		user.CreatedAt,
		user.Name,
		user.Email,
		user.Password,
		user.Address,
		user.PhoneNumber,
		user.IsSeller,
		user.IsDriver,
		user.Uuid,
		user.Nik,
	).Rows()

	if err != nil {
		return model.Users{}, err
	}

	defer rows.Close()

	var savedUser model.Users
	if rows.Next() {
		if err := r.db.ScanRows(rows, &savedUser); err != nil {
			return model.Users{}, err
		}
	}

	return savedUser, nil
}
