module github.com/harranali/examples/authentication

replace (
	github.com/harranali/examples/authentication/config => ./config
	github.com/harranali/examples/authentication/http => ./http
	github.com/harranali/examples/authentication/http/middlewares => ./http/middlewares
	github.com/harranali/examples/authentication/models => ./models
)

go 1.16

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gocondor/core v0.0.0-20210422202316-a6c1295d69b4
	github.com/joho/godotenv v1.3.0
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	gorm.io/gorm v1.21.6
)
