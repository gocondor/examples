// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"github.com/gin-gonic/gin"
)

// HomeShow to show home page
func HomeGet(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "welcome to the Todo example!",
	})
}
