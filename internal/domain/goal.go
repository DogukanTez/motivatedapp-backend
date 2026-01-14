package domain

import "time"

type GoalType string

const (
	GoalTypeDo    GoalType = "do"
	GoalTypeAvoid GoalType = "avoid"
)

type Goal struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	Title     string     `json:"title"`
	Type      GoalType   `json:"type"`
	StartDate string     `json:"start_date"` // "YYYY-MM-DD"
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
