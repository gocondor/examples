module github.com/gocondor/examples/authentication

replace (
	github.com/gocondor/core => C:/Users/harran/go/src/github.com/gocondor/core

	github.com/gocondor/examples/authentication/config => ./config
	github.com/gocondor/examples/authentication/http => ./http
	github.com/gocondor/examples/authentication/http/middlewares => ./http/middlewares
	github.com/gocondor/examples/authentication/integrations => ./integrations
	github.com/gocondor/examples/authentication/models => ./models
)

go 1.16

require (
	github.com/gin-gonic/autotls v0.0.3 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/gocondor/core v0.0.0-20210407100326-7c4af7b4c5a0
	github.com/google/uuid v1.2.0
	github.com/joho/godotenv v1.3.0
	golang.org/x/crypto v0.0.0-20210415154028-4f45737414dc
	golang.org/x/net v0.0.0-20210415231046-e915ea6b2b7d // indirect
	golang.org/x/sys v0.0.0-20210415045647-66c3f260301c // indirect
	gorm.io/gorm v1.21.6
)
