package input

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterUserRequest struct {
	Name        string  `json:"name" validate:"required,valid_string"`
	Email       string  `json:"email" binding:"required,email"`
	Password    string  `json:"password" binding:"required"`
	PhoneNumber string  `json:"phone_number" validate:"required,valid_phone_number"`
	Address     string  `json:"address" validate:"required,valid_address"`
	IsSeller    bool    `json:"is_seller"`
	IsDriver    bool    `json:"is_driver"`
	Nik         string  `json:"nik" validate:"required,valid_nik"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}
