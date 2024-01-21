package formatter

import (
	"github.com/google/uuid"
	"github.com/jajapp/domain/model"
)

type UserFormatter struct {
	Uuid    uuid.UUID `json:"uuid"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Address string    `json:"address"`
	Token   string    `json:"token"`
}

func FormatLoginUser(user model.Users, token string) UserFormatter {
	formatter := UserFormatter{
		Uuid:    user.Uuid,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Token:   token,
	}

	return formatter
}
