package input

type User struct {
	Name string `json:"name" binding:"required,alphanum"`
	Age  int    `json:"age" binding:"required,alphanum,min=18"`
}

type SigninData struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}
