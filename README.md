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

### 2- Authentication
A registration an authentication api with the following routes: 
```go
	router.Get("/", Auth, handlers.HomeShow)
	router.Post("/signup", handlers.UsersSignup)
	router.Post("/signin", handlers.UsersSignin)
	router.Get("/signout", handlers.UsersSignout)
```
