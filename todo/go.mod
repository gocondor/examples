module github.com/gocondor/examples/todo

replace (
	github.com/gocondor/examples/todo/config => ./config
	github.com/gocondor/examples/todo/http => ./http
	github.com/gocondor/examples/todo/http/middlewares => ./http/middlewares
	github.com/gocondor/examples/todo/http/handlers => ./http/handlers
	github.com/gocondor/examples/todo/models => ./models
)

go 1.16

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gocondor/core v0.0.0-20210422202316-a6c1295d69b4
	github.com/joho/godotenv v1.3.0
	gorm.io/gorm v1.21.6
)
