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
}

func FormatUser(user model.Users) UserFormatter {
	formatter := UserFormatter{
		Uuid:    user.Uuid,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
	}

	return formatter
}
