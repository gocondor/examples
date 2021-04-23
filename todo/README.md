# Todo example app
This is a todo API example using [GoCondor framework](https://gocondor.github.io)

## Routes 
Routes are located in `http/routes.go`
```go
	router.Get("/", handlers.HomeGet)
	router.Get("/todos", handlers.TodosList)
	router.Post("/todos", handlers.TodosCreate)
	router.Get("/todos/:id", handlers.TodosShow)
	router.Delete("/todos/:id", handlers.TodosDelete)
```

# Running the app 
1. Clone the `examples` repository and next `cd` into the Authentication app
2. Add `mysql` connection information to the `.env` file
3. Start your `mysql` server and create a database with the name `todo`
4. Run `go mod tidy` from within the app dir
5. Start the app by running the following command `go run main.go`
6. Open up your browser and navigate to `localhost:8000`
