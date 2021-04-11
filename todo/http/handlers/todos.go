// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocondor/core"
	"github.com/gocondor/examples/todo/models"
	"gorm.io/gorm"
)

// TodosList shows all todos
func TodosList(c *gin.Context) {
	// Get the database var from context
	db := c.MustGet(core.GORM).(*gorm.DB)

	var todos []models.Todo
	// Fetch all todos from the database
	result := db.Find(&todos)
	// Handle error
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(200, gin.H{
		"data": todos,
	})
}

// TodosCreate create a todo
func TodosCreate(c *gin.Context) {
	// Get the database var from context
	db := c.MustGet(core.GORM).(*gorm.DB)

	var todo models.Todo
	// Bind the input to the model
	err := c.ShouldBind(&todo)
	// Handle error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Store the record
	result := db.Create(&todo)
	// Handle Error
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(200, gin.H{
		"data": todo,
	})
}

// TodosList shows all todos
func TodosShow(c *gin.Context) {
	// Get the database var from context
	db := c.MustGet(core.GORM).(*gorm.DB)
	id := c.Param("id")

	var todo models.Todo
	// Get the record by id
	result := db.First(&todo, id)
	// Handle error
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})
}

// TodosDelete delete a todo
func TodosDelete(c *gin.Context) {
	// Get the database var from context
	db := c.MustGet(core.GORM).(*gorm.DB)
	id := c.Param("id")

	var todo models.Todo
	// Delete the record by id
	result := db.Delete(&todo, id)
	// Handle error
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted successfully",
	})
}
