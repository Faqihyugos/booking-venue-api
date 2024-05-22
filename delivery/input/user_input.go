package user

type RegisterUserInput struct {
	Email       string `json:"email" validate:"required,email"`
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required,min=6"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Fullname    string `json:"fullname" validate:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UpdateUserInput struct {
	ID          int
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Fullname    string `json:"fullname" validate:"required"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Error       error
}
