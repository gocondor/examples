# Examples
This repository contains example apps created using [GoCondor framework](https://gocondor.github.io)

## what are the examples?
### 1- Todo API
A todo api with the below routes:
```go
	router.Get("/", handlers.HomeGet)
	router.Get("/todos", handlers.TodosList)
	router.Post("/todos", handlers.TodosCreate)
	router.Get("/todos/:id", handlers.TodosShow)
	router.Delete("/todos/:id", handlers.TodosDelete)
```
