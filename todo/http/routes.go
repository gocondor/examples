// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/gocondor/core/routing"
	"github.com/gocondor/examples/todo/http/handlers"
)

// RegisterRoutes to register your routes
func RegisterRoutes() {
	router := routing.Resolve()

	//Define your routes here
	router.Get("/", handlers.HomeGet)
	router.Get("/todos", handlers.TodosList)
	router.Post("/todos", handlers.TodosCreate)
	router.Get("/todos/:id", handlers.TodosShow)
	router.Delete("/todos/:id", handlers.TodosDelete)

}
