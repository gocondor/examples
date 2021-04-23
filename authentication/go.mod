module github.com/gocondor/examples/authentication

replace (
	github.com/gocondor/examples/authentication/config => ./config
	github.com/gocondor/examples/authentication/http => ./http
	github.com/gocondor/examples/authentication/http/handlers => ./http/handlers
	github.com/gocondor/examples/authentication/http/middlewares => ./http/middlewares
	github.com/gocondor/examples/authentication/models => ./models
)

go 1.16

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gocondor/core v0.0.0-20210422202316-a6c1295d69b4
	github.com/joho/godotenv v1.3.0
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	gorm.io/gorm v1.21.6
)
