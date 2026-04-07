package model

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rana-touseef11/go-chi-postgresql/internal/dto"
	"github.com/rana-touseef11/go-chi-postgresql/pkg/constant"
)

type User struct {
	ID pgtype.UUID `json:"id"`

	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`

	Address *string              `json:"address,omitempty"`
	Status  *constant.UserStatus `json:"status,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func (u *User) UpdateUserFromDTO(req dto.UpdateUserRequest) {
	if req.Name != nil {
		u.Name = *req.Name
	}
	if req.Email != nil {
		u.Email = *req.Email
	}
	if req.Phone != nil {
		u.Phone = *req.Phone
	}
	if req.Address != nil {
		u.Address = req.Address
	}
}
