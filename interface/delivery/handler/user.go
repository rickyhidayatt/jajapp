package handler

import (
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
	var input input.LoginUser

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

	formatter := formatter.FormatLoginUser(loginUser, token)

	response := utils.ApiResponse("successfully login", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
