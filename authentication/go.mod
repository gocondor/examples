module github.com/gocondor/examples/authentication

replace (
	github.com/gocondor/examples/authentication/config => ./config
	github.com/gocondor/examples/authentication/http => ./http
	github.com/gocondor/examples/authentication/http/middlewares => ./http/middlewares
	github.com/gocondor/examples/authentication/integrations => ./integrations
	github.com/gocondor/examples/authentication/models => ./models
)

go 1.16

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gocondor/core v0.0.0-20210407100326-7c4af7b4c5a0
	github.com/joho/godotenv v1.3.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	golang.org/x/crypto v0.0.0-20210415154028-4f45737414dc // indirect
	golang.org/x/net v0.0.0-20210415231046-e915ea6b2b7d // indirect
	golang.org/x/sys v0.0.0-20210415045647-66c3f260301c // indirect
	golang.org/x/term v0.0.0-20210406210042-72f3dc4e9b72 // indirect
	google.golang.org/protobuf v1.23.0
	gorm.io/gorm v1.21.6
)
