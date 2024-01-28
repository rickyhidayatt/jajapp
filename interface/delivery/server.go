package delivery

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jajapp/database"
	"github.com/jajapp/domain/repository"
	"github.com/jajapp/domain/service"
	"github.com/jajapp/interface/delivery/auth"
	"github.com/jajapp/interface/delivery/handler"
)

func Run() {

	db, _ := database.Connect()
	repo := repository.NewUserRepository(db)

	//USER
	user := service.NewUserUsecase(repo)
	auth := auth.NewService()

	userHandler := handler.NewUserHandler(user, auth)

	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("api/v1")

	// user Endpoint
	api.POST("/user/login", userHandler.LoginUser)
	api.POST("/user/register", userHandler.RegisterUser)
	api.GET("/")
	router.Run()
}
