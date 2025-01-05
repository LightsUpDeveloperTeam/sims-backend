package authentication

import (
	"time"
)

type Schools struct {
	ID                    int64     `gorm:"primaryKey" json:"id"`
	Name                  string    `json:"name"`
	Address               string    `json:"address"`
	ContactEmail          string    `gorm:"unique" json:"contact_email"`
	ContactPhone          string    `json:"contact_phone"`
	LogoURL               string    `json:"logo_url"`
	SubscriptionStatus    string    `json:"subscription_status"`
	SubscriptionPlan      string    `json:"subscription_plan"`
	SubscriptionExpiryDate time.Time `json:"subscription_expiry_date"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

type User struct {
	ID           int64     `gorm:"primaryKey" json:"id"`
	SchoolID     int64     `gorm:"index" json:"school_id"` 
	Email        string    `gorm:"unique" json:"email"`
	Phone        string    `json:"phone"`
	SocialLogins string    `gorm:"type:jsonb" json:"social_logins"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type OTP struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	UserID    int64     `gorm:"index" json:"user_id"` 
	OTPCode   string    `json:"otp_code"`
	Method    string    `json:"method"`
	ExpiresAt time.Time `json:"expires_at"`
	IsUsed    bool      `json:"is_used"`
	CreatedAt time.Time `json:"created_at"`
}
