package domain

import "time"

type GoalType string

const (
	GoalTypeDo    GoalType = "do"
	GoalTypeAvoid GoalType = "avoid"
)

type Goal struct {
	ID        int64      `json:"id"`
	UserID    int64      `json:"user_id"`
	Title     string     `json:"title"`
	Type      GoalType   `json:"type"`
	StartDate time.Time  `json:"start_date"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
