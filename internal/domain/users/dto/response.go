package dto

type RegisterUserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type RegistrationResponse struct {
	Success bool                 `json:"success"`
	Message string               `json:"message"`
	Data    RegisterUserResponse `json:"data"`
}
