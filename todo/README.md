# Todo example app
This a todo API example using [GoCondor framework](gocondor.github.io)

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
First add your database `mysql` connection information to the `.env` file, then start the app by running the following command 
```bash
go run main.go
```
