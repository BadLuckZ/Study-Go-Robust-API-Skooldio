package main

import (
	"net/http"

	"github.com/BadLuckZ/Study-Go-Robust-API-Skooldio/auth"
	"github.com/BadLuckZ/Study-Go-Robust-API-Skooldio/todo"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Connect to database
	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Create new table: Todo
	db.AutoMigrate(&todo.Todo{})

	r := gin.Default()

	// Initialize Todo handler
	handler := todo.NewTodoHandler(db)

	// GET /ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "pong",
		})
	})

	// GET /token
	r.GET("/token", auth.AccessToken)

	// GET /todos
	r.GET("/todos", handler.GetTasks)

	// POST /todos
	r.POST("/todos", handler.NewTask)

	r.Run()
}
