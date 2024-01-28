package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jajapp/domain/public/formatter"
	"github.com/jajapp/domain/public/input"
	"github.com/jajapp/domain/service"
	"github.com/jajapp/interface/delivery/auth"
	"github.com/jajapp/utils"
)

type userHandler struct {
	userService service.UserServiceInterface
	authService auth.Service
}

func NewUserHandler(us service.UserServiceInterface, au auth.Service) *userHandler {
	return &userHandler{
		userService: us,
		authService: au,
	}
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input input.LoginUserRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := utils.FormatValidatorError(err)
		errorsMessage := gin.H{"error": errors}

		response := utils.ApiResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loginUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := utils.ApiResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loginUser.Uuid)
	if err != nil {
		response := utils.ApiResponse("failed login your account", http.StatusBadRequest, "auth login error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := formatter.UserResponse(loginUser, token)

	response := utils.ApiResponse("successfully login", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input input.RegisterUserRequest

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := utils.ApiResponse("Invalid request payload", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userRegister, err := h.userService.Register(input)
	if err != nil {
		// Log the error
		log.Printf("Failed to register user: %v", err)

		response := utils.ApiResponse("Failed to register user", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	token, err := h.authService.GenerateToken(userRegister.Uuid)
	if err != nil {
		response := utils.ApiResponse("failed login your account", http.StatusBadRequest, "auth login error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := formatter.UserResponse(*userRegister, token)

	response := utils.ApiResponse("Successfully registered user", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
