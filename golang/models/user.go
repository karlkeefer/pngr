package models

import (
	"encoding/json"
	"time"
)

type User struct {
	ID           int64
	Name         *string // nullable
	Email        string
	Salt         string
	Pass         string
	Status       Status
	Verification string
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Status is like a role, including unverified
type UserStatus int

const (
	UserStatusDisabled   Status = -1
	UserStatusUnverified        = 0
	UserStatusActive            = 1
	UserStatusAdmin             = 10
)

// MarshalJSON here protects "private" fields from ever being sent *out*
// it also makes Name return "" instead of null
func (u User) MarshalJSON() ([]byte, error) {
	var tmp struct {
		ID        int64      `json:"id"`
		Name      string     `json:"name"`
		Email     string     `json:"email,omitempty"`
		Status    Status     `json:"status"`
		CreatedAt *time.Time `json:"created_at,omitempty"`
	}

	tmp.ID = u.ID

	// pick a name
	if u.Name != nil {
		tmp.Name = *u.Name
	} else {
		tmp.Name = ""
	}

	tmp.Email = u.Email
	tmp.Status = u.Status
	if !u.CreatedAt.IsZero() {
		tmp.CreatedAt = &u.CreatedAt
	}
	return json.Marshal(&tmp)
}
