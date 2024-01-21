package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jajapp/config"
)

type Service interface {
	GenerateToken(uuid uuid.UUID) (string, error)
	ValidateToken(encodeToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (j *jwtService) GenerateToken(uuid uuid.UUID) (string, error) {
	claim := jwt.MapClaims{}
	claim["uuid"] = uuid

	config.ReloadEnv()
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedMethod, err := token.SignedString(secretKey)
	if err != nil {
		return signedMethod, err
	}

	return signedMethod, nil
}

func (j *jwtService) ValidateToken(encodeToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodeToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		config.ReloadEnv()
		secretKey := []byte(os.Getenv("SECRET_KEY"))

		return secretKey, nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
