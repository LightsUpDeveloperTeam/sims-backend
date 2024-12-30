package authentication

import "time"

type Tenant struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID            int64     `json:"id"`
	TenantID      int64     `json:"tenant_id"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	SocialLogins  string    `json:"social_logins"` 
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type OTP struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	OTPCode   string    `json:"otp_code"`
	Method    string    `json:"method"`
	ExpiresAt time.Time `json:"expires_at"`
	IsUsed    bool      `json:"is_used"`
	CreatedAt time.Time `json:"created_at"`
}

type Role struct {
	ID        int64     `json:"id"`
	TenantID  int64     `json:"tenant_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Permission struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RolePermission struct {
	RoleID    int64     `json:"role_id"`
	PermissionID int64     `json:"permission_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}