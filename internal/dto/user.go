package dto

type UserLoginRequest struct {
	Eamil string `json:"email" validate:"required,email"`
}

type UserBase struct {
	Name    string  `json:"name" validate:"required,min=2,max=50"`
	Email   string  `json:"email" validate:"required,email"`
	Phone   string  `json:"phone" validate:"required"`
	Address *string `json:"address,omitempty" validate:"omitempty"` // * make this field optional
}

type CreateUserRequest struct {
	UserBase
}

type UpdateUserRequest struct {
	Name    *string `json:"name,omitempty" validate:"required,min=2,max=50"`
	Email   *string `json:"email,omitempty" validate:"required,email"`
	Phone   *string `json:"phone,omitempty" validate:"required"`
	Address *string `json:"address,omitempty"` // * make this field optional
}
