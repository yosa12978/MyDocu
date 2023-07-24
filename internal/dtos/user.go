package dtos

type UserDTO struct {
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role"`
}

type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email,omitempty"`
}
