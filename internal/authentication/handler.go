package authentication

import (
	"sims-backend/internal/utils"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthHandler struct {
	Service *AuthService
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	repo := NewAuthRepository(db)
	service := NewAuthService(repo)
	return &AuthHandler{Service: service}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req struct {
		Email string `json:"email"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}

	err := h.Service.GenerateOTP(req.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.CreateResponse(
			"ERROR",
			err.Error(),
			nil,
			nil, nil, nil, nil,
		))
	}

	// accessToken, err := generateAccessToken(req.Email)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
	// 		"ERROR",
	// 		"Failed to generate access token",
	// 		nil,
	// 		nil, nil, nil, nil,
	// 	))
	// }

	// refreshToken, err := generateRefreshToken(req.Email)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
	// 		"ERROR",
	// 		"Failed to generate refresh token",
	// 		nil,
	// 		nil, nil, nil, nil,
	// 	))
	// }

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"OTP sent to email",
		nil,
		nil, nil, nil, nil,
	))
}


func (h *AuthHandler) VerifyOTP(c *fiber.Ctx) error {
	var req struct {
		Email   string `json:"email"`
		OTPCode string `json:"otp_code"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid request payload", nil, nil, nil, nil, nil,
		))
	}

	token, err := h.Service.VerifyOTP(req.Email, req.OTPCode)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil,
		))
	}

	
	refreshToken, err := generateRefreshToken(req.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to generate refresh token",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"OTP verified successfully",
		map[string]string{"accessToken": token, "refreshToken": refreshToken},
		nil, nil, nil, nil,
	))
}

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		))
	}

	_, claims, err := validateToken(req.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid or expired refresh token",
			nil,
			nil, nil, nil, nil,
		))
	}

	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "refresh" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid token type",
			nil,
			nil, nil, nil, nil,
		))
	}

	email, ok := claims["email"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.CreateResponse(
			"ERROR",
			"Invalid token payload",
			nil,
			nil, nil, nil, nil,
		))
	}

	newAccessToken, err := generateAccessToken(email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to generate access token",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Access token refreshed successfully",
		map[string]string{"NewAccessToken": newAccessToken},
		nil, nil, nil, nil,
	))
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	err := redisClient.Set(ctx, tokenString, "blacklisted", time.Hour).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.CreateResponse(
			"ERROR",
			"Failed to blacklist token",
			nil,
			nil, nil, nil, nil,
		))
	}

	return c.Status(fiber.StatusOK).JSON(utils.CreateResponse(
		"SUCCESS",
		"Logged out successfully",
		nil,
		nil, nil, nil, nil,
	))
}


