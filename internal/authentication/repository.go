package authentication

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type AuthRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{DB: db}
}

func (repo *AuthRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	if err := repo.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (repo *AuthRepository) SaveOTP(otp OTP) error {
	return repo.DB.Create(&otp).Error
}

func (repo *AuthRepository) GetOTP(userID int64, otpCode string) (*OTP, error) {
	var otp OTP
	if err := repo.DB.Where("user_id = ? AND otp_code = ? AND is_used = false AND expires_at > ?", userID, otpCode, time.Now()).First(&otp).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("otp not found or expired")
		}
		return nil, err
	}
	return &otp, nil
}

func (repo *AuthRepository) MarkOTPAsUsed(otpID int64) error {
	return repo.DB.Model(&OTP{}).Where("id = ?", otpID).Update("is_used", true).Error
}
