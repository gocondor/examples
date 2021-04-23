package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocondor/examples/todo/models"
)

// TodosList shows all todos
func TodosList(c *gin.Context) {
	var todos []models.Todo
	// Fetch all todos from the database
	result := DB.Find(&todos)
	// Handle error
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": todos,
	})
}

// TodosCreate create a todo
func TodosCreate(c *gin.Context) {
	var todo models.Todo
	// Bind the input to the model
	err := c.ShouldBind(&todo)
	// Handle error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Store the record
	result := DB.Create(&todo)
	// Handle Error
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": todo,
	})
}

// TodosList shows all todos
func TodosShow(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo
	// Get the record by id
	result := DB.First(&todo, id)
	// Handle error
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}

// TodosDelete delete a todo
func TodosDelete(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo
	// Delete the record by id
	result := DB.Delete(&todo, id)
	// Handle error
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted successfully",
	})
}
