package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jajapp/domain/model"
	"github.com/jajapp/domain/public/input"
	"github.com/jajapp/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Login(input input.LoginUserRequest) (model.Users, error)
	Register(input input.RegisterUserRequest) (*model.Users, error)
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

func (s *userService) Register(input input.RegisterUserRequest) (*model.Users, error) {
	checkEmail, err := s.userRepo.FindEmail(input.Email)

	if checkEmail.Email != "" {
		return nil, errors.New("email already exists")
	}

	if input.IsSeller == true && input.IsDriver == true {
		return nil, errors.New("You cannot be both a seller and a driver")
	}

	userID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return nil, err
	}

	var user = model.Users{
		CreatedAt:   time.Now(),
		Uuid:        userID,
		Name:        input.Name,
		Email:       input.Email,
		Password:    string(passwordHash),
		Address:     input.Address,
		PhoneNumber: input.PhoneNumber,
		IsSeller:    input.IsSeller,
		IsDriver:    input.IsDriver,
		Nik:         input.Nik,
	}

	newUser, err := s.userRepo.SaveUser(user)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
