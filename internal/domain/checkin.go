package domain

import (
	"time"
)

type CheckinStatus string

const (
	CheckinSuccess CheckinStatus = "success"
	CheckinFail    CheckinStatus = "fail"
)

type Checkin struct {
	ID int `json:"id"`

	GoalID int `json:"goal_id"`

	CheckinDate time.Time     `json:"checkin_date"` // YYYY-MM-DD
	Status      CheckinStatus `json:"status"`

	CreatedAt time.Time `json:"created_at"`
}
