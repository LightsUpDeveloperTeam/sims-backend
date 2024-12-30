package authentication

import (
	"database/sql"
	"errors"
	"log"
	"sims-backend/internal/database"
	"time"
)

type AuthRepository struct {
	DB *sql.DB
}

func NewAuthRepository(db database.Service) *AuthRepository {
	return &AuthRepository{DB: db.GetDB()}
}

func (repo *AuthRepository) GetUserByEmail(email string) (*User, error) {
	log.Printf("Searching user by email: %s", email)

	var user User
	query := `SELECT id, tenant_id, email, phone, social_logins, created_at, updated_at FROM users WHERE email = $1`
	err := repo.DB.QueryRow(query, email).Scan(
		&user.ID, &user.TenantID, &user.Email, &user.Phone, &user.SocialLogins, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("No user found with email: %s", email)
			return nil, errors.New("user not found")
		}
		log.Printf("Error querying user: %v", err)
		return nil, err
	}
	return &user, nil
}

func (repo *AuthRepository) SaveOTP(otp OTP) error {
	log.Printf("Saving OTP for user ID: %d", otp.UserID)

	query := `INSERT INTO otps (user_id, otp_code, method, expires_at, is_used, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := repo.DB.Exec(query, otp.UserID, otp.OTPCode, otp.Method, otp.ExpiresAt, otp.IsUsed, otp.CreatedAt)
	if err != nil {
		log.Printf("Error saving OTP: %v", err)
	}
	return err
}

func (repo *AuthRepository) GetOTP(userID int64, otpCode string) (*OTP, error) {
	log.Printf("Fetching OTP for user ID: %d", userID)

	var otp OTP
	query := `SELECT id, user_id, otp_code, method, expires_at, is_used, created_at FROM otps WHERE user_id = $1 AND otp_code = $2 AND is_used = false AND expires_at > $3`
	err := repo.DB.QueryRow(query, userID, otpCode, time.Now()).Scan(
		&otp.ID, &otp.UserID, &otp.OTPCode, &otp.Method, &otp.ExpiresAt, &otp.IsUsed, &otp.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("OTP not found or expired for user ID: %d", userID)
			return nil, errors.New("otp not found or expired")
		}
		log.Printf("Error fetching OTP: %v", err)
		return nil, err
	}
	return &otp, nil
}

func (repo *AuthRepository) MarkOTPAsUsed(otpID int64) error {
	log.Printf("Marking OTP as used for OTP ID: %d", otpID)

	query := `UPDATE otps SET is_used = true WHERE id = $1`
	_, err := repo.DB.Exec(query, otpID)
	if err != nil {
		log.Printf("Error marking OTP as used: %v", err)
	}
	return err
}
