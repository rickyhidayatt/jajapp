package input

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterUserRequest struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber int    `json:"phone_number" validate:"required"`
	Address     string `json:"address" validate:"required"`
	IsSeller    bool   `json:"is_seller"`
	IsDriver    bool   `json:"is_driver"`
}
