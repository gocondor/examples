// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"time"

	"gorm.io/gorm"
)

// Todo model
type Todo struct {
	ID        uint   `gorm:"primarykey"`
	Title     string `json:"title" form:"title" binding:"required"`
	Body      string `json:"body" form:"body" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
