package request

type UserRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=10"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=10,containsany=$@%*!"`
	Age      int    `json:"age" binding:"required,min=14,max=110"`
}
