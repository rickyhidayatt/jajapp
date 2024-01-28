package formatter

import (
	"time"

	"github.com/google/uuid"
	"github.com/jajapp/domain/model"
)

type userResponse struct {
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"register_date"`
}

func UserResponse(user model.Users, token string) userResponse {
	formatter := userResponse{
		Uuid:      user.Uuid,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Token:     token,
		CreatedAt: user.CreatedAt,
	}

	return formatter
}
