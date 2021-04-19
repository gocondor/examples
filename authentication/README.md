# Authentication example app
This is a registration and authentication API example using [GoCondor framework](https://gocondor.github.io)

## Routes 
Routes are located in `http/routes.go`
```go
	router.Get("/", Auth, handlers.HomeShow)
	router.Post("/signup", handlers.UsersSignup)
	router.Post("/signin", handlers.UsersSignin)
	router.Get("/signout", handlers.UsersSignout)
```

# Running the app 
First clone the `examples` repository, next `cd` into the Authentication app, add your database `mysql` and `redis` connection information to the `.env` file, then start the app by running the following command 
```bash
go run main.go
```
