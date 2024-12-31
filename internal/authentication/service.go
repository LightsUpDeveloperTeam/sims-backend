package authentication

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	Repo *AuthRepository
}

func NewAuthService(repo *AuthRepository) *AuthService {
	return &AuthService{Repo: repo}
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET_KEY")) 

func (s *AuthService) GenerateOTP(email string) error {
	log.Printf("Generating OTP for email: %s", email)

	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		log.Printf("Error finding user: %v", err)
		return err
	}

	otpCode := generateRandomOTP()
	otp := OTP{
		UserID:    user.ID,
		OTPCode:   otpCode,
		Method:    "email",
		ExpiresAt: time.Now().Add(5 * time.Minute),
		IsUsed:    false,
		CreatedAt: time.Now(),
	}

	err = s.Repo.SaveOTP(otp)
	if err != nil {
		log.Printf("Error saving OTP: %v", err)
		return err
	}

	// Mock sending OTP via email
	log.Printf("OTP %s sent to email %s", otpCode, user.Email)
	return nil
}

func (s *AuthService) VerifyOTP(email, otpCode string) (string, error) {
	log.Printf("Verifying OTP for email: %s", email)

	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		log.Printf("Error finding user: %v", err)
		return "", errors.New("user not found")
	}

	otp, err := s.Repo.GetOTP(user.ID, otpCode)
	if err != nil {
		log.Printf("Error verifying OTP: %v", err)
		return "", errors.New("invalid or expired OTP")
	}

	err = s.Repo.MarkOTPAsUsed(otp.ID)
	if err != nil {
		log.Printf("Error marking OTP as used: %v", err)
		return "", errors.New("failed to mark OTP as used")
	}

	token, err := generateAccessToken(user.Email)
	if err != nil {
		log.Printf("Error generating access token: %v", err)
		return "", errors.New("failed to generate access token")
	}

	log.Printf("OTP verified successfully, token generated: %s", token)
	return token, nil
}


func generateRandomOTP() string {
	number := make([]byte, 6)
	_, err := rand.Read(number)
	if err != nil {
		log.Printf("Error generating random OTP: %v", err)
		return "000000"
	}
	return fmt.Sprintf("%06d", int(number[0])*10000+int(number[1])*1000+int(number[2])*100+int(number[3])*10+int(number[4]))
}

func generateAccessToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,                     
		"exp":   time.Now().Add(1 * time.Hour).Unix(), 
		"iat":   time.Now().Unix(),        
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}