package formatter

import (
	"time"

	"github.com/google/uuid"
	"github.com/jajapp/domain/model"
)

type userResponse struct {
	Uuid        uuid.UUID     `json:"uuid"`
	Name        string        `json:"name"`
	Email       string        `json:"email"`
	PhoneNumber string        `json:"phone_number"`
	Location    AddressDetail `json:"location"`
	Token       string        `json:"token"`
	CreatedAt   time.Time     `json:"register_date"`
	IsSeller    bool          `json:"is_seller"`
	IsDriver    bool          `json:"is_driver"`
	IsVerified  bool          `json:"is_verified"`
	Nik         string        `json:"nik"`
}

type AddressDetail struct {
	DetailAddress string `json:"detail_address"`
	GeoLocation   geo    `json:"geo_location"`
}

type geo struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func UserResponse(user model.Users, token string) userResponse {
	geo := geo{
		Latitude:  user.Latitude,
		Longitude: user.Longitude,
	}

	address := AddressDetail{
		DetailAddress: user.Address,
		GeoLocation:   geo,
	}

	formatter := userResponse{
		Uuid:        user.Uuid,
		Name:        user.Name,
		Email:       user.Email,
		Location:    address,
		PhoneNumber: user.PhoneNumber,
		Token:       token,
		IsSeller:    user.IsSeller,
		IsDriver:    user.IsDriver,
		CreatedAt:   user.CreatedAt,
		Nik:         user.Nik,
	}

	return formatter
}
