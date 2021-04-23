package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocondor/examples/authentication/http/input"
	"github.com/gocondor/examples/authentication/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UsersSignup to create new accounts
func UsersSignup(c *gin.Context) {
	// bind the input to the user's model
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}
	// check if there is a record with the given email
	res := DB.Where("email = ?", user.Email).First(&models.User{})
	if res.Error == nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "user already signed up",
		})
		return
	}
	//hash the passowrd
	hahsedPWD, err := hashPassword(user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	//use the hashed password
	user.Password = hahsedPWD
	// create the db record
	res = DB.Create(&user)
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

func UsersSignin(c *gin.Context) {
	// validate and bind user input
	var signinData input.SigninData
	if err := c.ShouldBind(&signinData); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
	}
	// get the user record by email from db
	var user models.User
	result := DB.Where("email = ?", signinData.Email).First(&user)
	// check if the record not found
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "wrong credentials",
		})
		return
	}
	// handle database error incase there is any
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "someting went wrong",
		})
		return
	}
	//compare the password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(signinData.Password))
	if err != nil {
		// wrong password
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "wrong credentials",
		})
		return
	}
	// prepare jwt token payload
	tokenPayload := map[string]string{
		"userId": strconv.FormatUint(uint64(user.ID), 10),
	}
	// generate the jwt token
	token, err := JWT.CreateToken(tokenPayload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
	}

	// store the token in redis
	tokenRedisKey := fmt.Sprintf("%s-token", tokenPayload["userId"])
	Cache.Set(tokenRedisKey, strconv.FormatUint(uint64(user.ID), 10))

	// prepare the refresh token payload
	refreshTokenPayload := map[string]string{
		"userId": strconv.FormatUint(uint64(user.ID), 10),
	}
	// generate the token
	refreshToken, err := JWT.CreateRefreshToken(refreshTokenPayload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
	}
	// store the refresh token in redis
	refreshRedisKey := fmt.Sprintf("%s-refresh", refreshTokenPayload["userId"])
	Cache.Set(refreshRedisKey, strconv.FormatUint(uint64(user.ID), 10))
	// render response
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]string{
			"token":        token,
			"refreshToken": refreshToken,
		},
	})
}

func UsersSignout(c *gin.Context) {
	// extract the token
	token, err := JWT.ExtractToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request data",
		})
		return
	}
	// decode the token
	payload, err := JWT.DecodeToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad request data",
		})
		return
	}

	// delete the token redis entry
	tokenRedisKey := fmt.Sprintf("%s-token", payload["userId"])
	err = Cache.Delete(tokenRedisKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}
	// delete the refresh token redis entry
	refreshRedisKey := fmt.Sprintf("%s-refresh", payload["userId"])
	err = Cache.Delete(refreshRedisKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "signed out successfully",
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
