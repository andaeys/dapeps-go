// main.go
package main

import (
	"dapeps-go/db"
	"dapeps-go/delivery/http"
	"dapeps-go/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//init DB
	db.InitDB()
	dbConnection := db.GetDB()

	//provide the repo
	userRepository := user.NewUserRepository(dbConnection)

	//provide the service
	userService := &user.UserServiceImpl{UserRepository: &userRepository}

	//setup route handlers
	handlers := http.NewUserHandler(userService)
	handlers.SetupUserRoutes(r)

	r.Run(":8080")
}
