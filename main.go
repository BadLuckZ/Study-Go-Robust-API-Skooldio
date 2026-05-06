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

	// Create Protected Routes
	protected := r.Group("", auth.Protect([]byte("==signature==")))

	// Initialize Todo handler
	handler := todo.NewTodoHandler(db)

	// GET /ping
	// Public
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{
			"message": "pong",
		})
	})

	// GET /token
	// Public
	r.GET("/token", auth.AccessToken)

	// GET /todos
	// Public
	r.GET("/todos", handler.GetTasks)

	// POST /todos
	// Private
	protected.POST("/todos", handler.NewTask)

	r.Run()
}
