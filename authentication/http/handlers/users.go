// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocondor/core"
	"github.com/gocondor/examples/authentication/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// HomeShow to show home page
func UsersSignup(c *gin.Context) {
	// grab the db variable
	db := c.MustGet(core.GORM).(*gorm.DB)
	// bind the input to the user's model
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		fmt.Println("got error++++")
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}
	// check if there is a record with the given email
	res := db.Where("email = ?", user.Email).First(&models.User{})
	if res.Error == nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "user already signed up",
		})
		return
	}
	//encrypt the password filed
	hahsedPWD, err := hashPassword(user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	user.Password = hahsedPWD
	// create the db record
	res = db.Create(&user)
	if res.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": res.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "signup successfully",
	})
}

// hashPassword hashs passwords
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
