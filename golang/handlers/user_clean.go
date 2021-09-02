package handlers

import (
	"time"

	"github.com/karlkeefer/pngr/golang/db"
)

type UserForAPI struct {
	ID        int64         `json:"id"`
	Email     string        `json:"email,omitempty"`
	Status    db.UserStatus `json:"status"`
	CreatedAt time.Time     `json:"created_at,omitempty"`
}

// use this to avoid leaking sensitive data via API
func cleanUser(u *db.User) *UserForAPI {
	return &UserForAPI{
		ID:        u.ID,
		Email:     u.Email,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
	}
}
