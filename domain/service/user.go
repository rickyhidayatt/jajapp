package service

import (
	"github.com/jajapp/domain/model"
	"github.com/jajapp/domain/public/input"
	"github.com/jajapp/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Login(input input.LoginUser) (model.Users, error)
}

type userService struct {
	userRepo repository.UserRepositoryInterface
}

func NewUserUsecase(u repository.UserRepositoryInterface) UserServiceInterface {
	return &userService{userRepo: u}
}

func (s *userService) Login(input input.LoginUser) (model.Users, error) {

	user, err := s.userRepo.FindEmail(input.Email)
	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user, err
	}

	return user, nil
}
