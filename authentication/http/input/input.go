package input

// User represents request data with user information
type User struct {
	Name string `json:"name" binding:"exists,alphanum"`
	Age  int    `json:"age" binding:"exists,alphanum,min=18"`
}

type SigninData struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}
