// Copyright 2021 Harran Ali <harran.m@gmail.com>. All rights reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/gocondor/core/routing"
	"github.com/gocondor/examples/authentication/http/handlers"
	. "github.com/gocondor/examples/authentication/http/middlewares"
)

// RegisterRoutes to register your routes
func RegisterRoutes() {
	router := routing.Resolve()

	//Define your routes here
	router.Get("/", Auth, handlers.HomeShow)
	router.Post("/signup", handlers.UsersSignup)
	router.Post("/signin", handlers.UsersSignin)
	router.Get("/signout", handlers.UsersSignout)
}
