package domain

import "time"

type SubscriptionStatus string

const (
	SubscriptionStatusActive   SubscriptionStatus = "active"
	SubscriptionStatusExpired  SubscriptionStatus = "expired"
	SubscriptionStatusCanceled SubscriptionStatus = "canceled"
)

type Platform string

const (
	PlatformIOS     Platform = "ios"
	PlatformAndroid Platform = "android"
)

type Subscription struct {
	ID                    int64              `json:"id"`
	UserID                int64              `json:"user_id"`
	Platform              Platform           `json:"platform"`
	ProductID             string             `json:"product_id"`
	IsLifetime            bool               `json:"is_lifetime"`
	StartedAt             time.Time          `json:"started_at"`
	ExpiresAt             *time.Time         `json:"expires_at,omitempty"`
	Status                SubscriptionStatus `json:"status"`
	OriginalTransactionID *string            `json:"original_transaction_id,omitempty"`
	CreatedAt             time.Time          `json:"created_at"`
}
