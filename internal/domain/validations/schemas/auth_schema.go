package schemas

type UserRegisterSchema struct {
	FirstName            string `json:"first_name" binding:"required,min=3,max=50"`
	LastName             string `json:"last_name" binding:"required,min=3,max=100"`
	Email                string `json:"email" binding:"required,email,max=100"`
	Phone                string `json:"phone" binding:"required,min=11,max=14,numeric"`
	Password             string `json:"password" binding:"required,min=8,max=30"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
}

type UserLoginSchema struct {
	Email    string `json:"email" binding:"omitempty,email,max=100"`
	Phone    string `json:"phone" binding:"omitempty,min=11,max=14,numeric"`
	Password string `json:"password" binding:"required,min=8,max=30"`
}

type UserResendVerificationSchema struct {
	Email string `json:"email" binding:"omitempty,email,max=100"`
	Phone string `json:"phone" binding:"omitempty,min=11,max=14,numeric"`
}
