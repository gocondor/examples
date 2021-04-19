// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocondor/core"
	"github.com/gocondor/core/cache"
	"github.com/gocondor/core/jwtloader"
	"github.com/gocondor/examples/authentication/models"
	"gorm.io/gorm"
)

// MiddlewareExample is an example of a middleware gets executed before the request handler
var Auth gin.HandlerFunc = func(c *gin.Context) {
	// Get the jwt loader variable from context
	jwt := c.MustGet(core.JWT).(*jwtloader.JwtLoader)
	// Get the cache variable from context
	cache := c.MustGet(core.CACHE).(*cache.CacheEngine)
	// Get the db variable from context
	db := c.MustGet(core.GORM).(*gorm.DB)

	// extract the token
	token, err := jwt.ExtractToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	// validate the token
	_, err = jwt.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	// decode the token
	payload, err := jwt.DecodeToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}

	// check if user has a record in redis
	tokenRedisKey := fmt.Sprintf("%s-token", payload["userId"])
	res, err := cache.Get(tokenRedisKey)
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
	result := db.Find(&user, userId)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "unauthorized",
		})
		return
	}
	// Pass on to the next-in-chain
	c.Next()
}
