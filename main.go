package main

import (
	"9gag-api/handler"
	"9gag-api/model"
	"9gag-api/repository"
	"9gag-api/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}

	// TODO: Refactor db initialization
	model.DB, err = gorm.Open(sqlite.Open(os.Getenv("DB_SQLITE_DSN")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	model.DB.AutoMigrate(&model.User{})

	// TODO: move this to router package
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "pong"})
	})

	var userRepository = repository.NewUserRepository()
	var authService = service.NewAuthService(userRepository)
	var userHandler = handler.NewUserHandler(authService)

	r.POST("/signup", userHandler.Signup)
	r.POST("/signin", userHandler.Signin)

	r.Run()

}
