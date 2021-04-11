module github.com/gocondor/examples/todo

replace (
	github.com/gocondor/examples/todo/config => ./config
	github.com/gocondor/examples/todo/http => ./http
	github.com/gocondor/examples/todo/http/middlewares => ./http/middlewares
	github.com/gocondor/examples/todo/integrations => ./integrations
	github.com/gocondor/examples/todo/models => ./models
)

go 1.16

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gocondor/core v0.0.0-20210407100326-7c4af7b4c5a0
	github.com/joho/godotenv v1.3.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	gorm.io/gorm v1.21.6
)
