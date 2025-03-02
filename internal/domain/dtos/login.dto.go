package dtos

type LoginDto struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password,omitempty" binding:"required"`
    Token    string `json:"token,omitempty"`
}
