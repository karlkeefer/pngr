package db

import (
	"encoding/json"
	"time"
)

// MarshalJSON here protects "private" fields from ever being sent *out*
func (u User) MarshalJSON() ([]byte, error) {
	tmp := struct {
		ID        int64      `json:"id"`
		Email     string     `json:"email"`
		Status    UserStatus `json:"status"`
		CreatedAt time.Time  `json:"created_at"`
	}{
		u.ID,
		u.Email,
		u.Status,
		u.CreatedAt,
	}

	return json.Marshal(tmp)
}
