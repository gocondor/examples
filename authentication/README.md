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
1- Clone the `examples` repository and next `cd` into the Authentication app
2- Add `mysql` and `redis` connection information to the `.env` file
3- Start your `mysql` server and create a database with the name `authentication`
4- Start your `redis` server
5- Run `go mod tidy` from within the app dir
6- Start the app by running the following command `go run main.go`
7- Open up your browser and navigate to `localhost:8000`
