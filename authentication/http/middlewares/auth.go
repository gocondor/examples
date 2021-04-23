package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocondor/examples/authentication/models"
	"gorm.io/gorm"
)

// MiddlewareExample is an example of a middleware gets executed before the request handler
var Auth gin.HandlerFunc = func(c *gin.Context) {

	// extract the token
	token, err := JWT.ExtractToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	// validate the token
	_, err = JWT.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	// decode the token
	payload, err := JWT.DecodeToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	// check if user has a record in redis
	tokenRedisKey := fmt.Sprintf("%s-token", payload["userId"])
	res, err := Cache.Get(tokenRedisKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}
	// convert the user id to int64
	userId, _ := strconv.ParseInt(fmt.Sprintf("%s", res), 10, 64)
	// check db for user with given id
	var user models.User
	result := DB.Find(&user, userId)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}
	// Pass on to the next-in-chain
	c.Next()
}
