package todo

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// New table: Todos
type Todo struct {
	Title string `json:"text"`
	gorm.Model
}

// Todo Handler Datatype
type TodoHandler struct {
	db *gorm.DB
}

// Todo Handler Initialization
func NewTodoHandler(db *gorm.DB) *TodoHandler {
	return &TodoHandler{db: db}
}

// Create new todo item
func (t *TodoHandler) NewTask(c *gin.Context) {
	var todo Todo

	// Change JSON values from request to Todo's Object
	err := c.BindJSON(&todo)
	// From Todo Table -> require json must have "text" attribute

	// If there's an error while changing...
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Insert todo object into Todo Table
	r := t.db.Create(&todo)
	err = r.Error

	// If there's an error while creating
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Return successful status
	c.JSON(http.StatusCreated, gin.H{
		"id": todo.Model.ID,
	})
}

// Get all todo items
func (t *TodoHandler) GetTasks(c *gin.Context) {
	var todos []Todo

	r := t.db.Find(&todos)
	err := r.Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch todos",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    todos,
	})
}

func (t *TodoHandler) RemoveTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	r := t.db.Delete(&Todo{}, id)
	err = r.Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
