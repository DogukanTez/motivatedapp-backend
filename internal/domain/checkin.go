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
	ID int64 `json:"id"`

	GoalID int64 `json:"goal_id"`

	CheckinDate time.Time     `json:"checkin_date"` // YYYY-MM-DD
	Status      CheckinStatus `json:"status"`

	CreatedAt time.Time `json:"created_at"`
}
