// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	// init postgresql
	var err error
	db, err = gorm.Open("postgres", "postgres://anda:anda123@postgres:5432/anda_db?sslmode=disable")
	if err != nil {
		fmt.Println("error open postgre database")
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	r := gin.Default()

	// Define routes
	r.GET("/api/v1/users", getUsers)
	r.POST("/api/v1/users", createUser)
	r.GET("/api/v1/users/:id", getUserByID)
	r.PUT("/api/v1/users/:id", updateUser)
	r.DELETE("/api/v1/users/:id", deleteUser)

	// Start the server
	r.Run(":8080")
}

func getUsers(c *gin.Context) {
	var users []User
	db.Find(&users)

	c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindQuery(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the user into the database
	db.Create(&user)

	c.JSON(http.StatusOK, user)
}

func getUserByID(c *gin.Context) {
	// Get user by ID (dummy data for now)
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID",
		"id":      id,
	})
}

func updateUser(c *gin.Context) {
	// Update user by ID (dummy data for now)
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Update user",
		"id":      id,
	})
}

func deleteUser(c *gin.Context) {
	// Delete user by ID (dummy data for now)
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete user",
		"id":      id,
	})
}
