package dto

type CreateUserRequest struct {
	Name     string `json:"name"     validate:"required,min=1,max=32"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (r *CreateUserRequest) GetCustomMessagesForValidator() map[string]string {
	return map[string]string{
		// Name
		"CreateUserRequest.Name.required": "Имя является обязательным полем.",
		"CreateUserRequest.Name.min":      "Имя должно содержать минимум 1 символ.",
		"CreateUserRequest.Name.max":      "Имя должно содержать не более 32 символа.",
		// Email
		"CreateUserRequest.Email.required": "Адрес электронной почты является обязательным полем.",
		"CreateUserRequest.Email.email":    "Невалидный адрес электронной почты.",
		// Password
		"CreateUserRequest.Password.required": "Пароль является обязательным полем.",
		"CreateUserRequest.Password.min":      "Пароль должен содержать минимум 8 символов.",
	}
}
