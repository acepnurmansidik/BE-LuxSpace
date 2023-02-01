package user

type CreateUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
	Role     string `json:"role"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ActivateOtpInput struct {
	Otp string `uri:"otp" binding:"required"`
}
