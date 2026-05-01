package main

import (
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	// Connect to database
	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Create new table: User
	db.AutoMigrate(&User{})

	// Add new user
	res := db.Create(&User{Name: "Jack"})
	if res.Error != nil {
		panic(res.Error)
	}

	r := gin.Default()

	// GET /users : Get all users
	r.GET("/users", func(c *gin.Context) {
		var users []User
		db.Find(&users)
		c.JSON(200, users)
	})

	r.Run()
}
