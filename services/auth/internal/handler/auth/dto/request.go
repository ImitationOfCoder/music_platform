package dto

type LoginUserRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (r *LoginUserRequest) GetCustomMessagesForValidator() map[string]string {
	return map[string]string{
		// Email
		"LoginUserRequest.Email.required": "Адрес электронной почты является обязательным полем.",
		"LoginUserRequest.Email.email":    "Невалидный адрес электронной почты.",
		// Password
		"LoginUserRequest.Password.required": "Пароль является обязательным полем.",
		"LoginUserRequest.Password.min":      "Пароль должен содержать минимум 8 символов.",
	}
}
