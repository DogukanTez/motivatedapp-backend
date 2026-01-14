package domain

import "time"

type User struct {
	ID           int         `json:"id"`
	Email        *string     `json:"email,omitempty"`
	Password     *string     `json:"password,omitempty"`
	AuthProvider string      `json:"auth_provider"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    *time.Time  `json:"updated_at,omitempty"`
	UserDetail   *UserDetail `json:"user_detail,omitempty"`
}

type UserDetail struct {
	UserID           int        `json:"user_id"`
	DeviceID         *string    `json:"device_id,omitempty"`
	Platform         *string    `json:"platform,omitempty"`
	Locale           *string    `json:"locale,omitempty"`
	NotificationTime *string    `json:"notification_time,omitempty"` // HH:MM time format
	LastSeenAt       *time.Time `json:"last_seen_at,omitempty"`      //timestamptz
}
