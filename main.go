package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall"
	"time"

	"os/signal"

	"github.com/BadLuckZ/Study-Go-Robust-API-Skooldio/auth"
	"github.com/BadLuckZ/Study-Go-Robust-API-Skooldio/todo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"

	"gorm.io/gorm"
)

func main() {
	// Load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error in Environment Variables: %s\n", err)
	}

	// Connect to database
	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Create new table: Todo
	db.AutoMigrate(&todo.Todo{})

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",
	}

	config.AllowHeaders = []string{
		"Origin",
		"Authorization",
	}

	r.Use(cors.New(config))

	config.AllowHeaders = []string{}

	// Create Protected Routes
	protected := r.Group("",
		auth.RateLimit(),
		auth.Protect([]byte(os.Getenv("SIGN"))),
	)

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

	// ===================
	// Graceful Shutdown

	// Ctrl+C (SIGINT) to kill processes (SIGTERM)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	s := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	// Not go after this line until Ctrl+C
	<-ctx.Done()
	stop()
	fmt.Println("Shut down gracefully. Ctrl+C again to force")

	// Wait for 5 seconds before real shutdown
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = s.Shutdown(timeoutCtx)
	if err != nil {
		fmt.Println(err)
	}
	// ===================

}
