package service

import (
	"strconv"
	"time"

	"github.com/jajapp/domain/model"
	"github.com/jajapp/domain/public/input"
	"github.com/jajapp/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Login(input input.LoginUserRequest) (model.Users, error)
	Register(input input.RegisterUserRequest) (model.Users, error)
	UpdateProfile(input input.RegisterUserRequest) (model.Users, error)
}

type userService struct {
	userRepo repository.UserRepositoryInterface
}

func NewUserUsecase(u repository.UserRepositoryInterface) UserServiceInterface {
	return &userService{userRepo: u}
}

func (s *userService) Login(input input.LoginUserRequest) (model.Users, error) {

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

func (s *userService) Register(input input.RegisterUserRequest) (model.Users, error) {

	checkEmail, err := s.userRepo.CheckEmail(input.Email)
	if err != nil {
		return model.Users{}, err
	}

	if checkEmail != false {
		return model.Users{}, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return model.Users{}, err
	}

	var user = model.Users{
		CreatedAt:   time.Now(),
		Name:        input.Name,
		Email:       input.Email,
		Password:    string(passwordHash),
		Address:     input.Address,
		PhoneNumber: strconv.Itoa(input.PhoneNumber),
	}

	newUser, err := s.userRepo.SaveUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *userService) UpdateProfile(input input.RegisterUserRequest) (model.Users, error) {

	var user = model.Users{
		UpdatedAt:   time.Now(),
		Name:        input.Name,
		Email:       input.Email,
		Address:     input.Address,
		PhoneNumber: strconv.Itoa(input.PhoneNumber),
	}

	newUser, err := s.userRepo.Update(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
