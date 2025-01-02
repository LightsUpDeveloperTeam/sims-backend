package authentication

import (
	"crypto/rand"
	"errors"
	"fmt"
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
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
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

	return s.Repo.SaveOTP(otp)
}

func (s *AuthService) VerifyOTP(email, otpCode string) (string, error) {
	user, err := s.Repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	otp, err := s.Repo.GetOTP(user.ID, otpCode)
	if err != nil {
		return "", errors.New("invalid or expired OTP")
	}

	err = s.Repo.MarkOTPAsUsed(otp.ID)
	if err != nil {
		return "", errors.New("failed to mark OTP as used")
	}

	return generateAccessToken(user.Email)
}

func generateRefreshToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(7 * 24 * time.Hour).Unix(), 
		"iat":   time.Now().Unix(),
		"type":  "refresh", 
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}


func validateToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorClaimsInvalid)
	}

	return token, claims, nil
}

func generateRandomOTP() string {
	number := make([]byte, 6)
	_, _ = rand.Read(number)
	return fmt.Sprintf("%06d", int(number[0])%10*100000+int(number[1])%10*10000+int(number[2])%10*1000+int(number[3])%10*100+int(number[4])%10*10+int(number[5])%10)
}

func generateAccessToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(1 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
