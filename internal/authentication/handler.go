package authentication

import (
	"log"
	"sims-backend/internal/database"
	"sims-backend/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Service *AuthService
}

func NewAuthHandler(db database.Service) *AuthHandler {
	repo := NewAuthRepository(db)
	service := NewAuthService(repo)
	return &AuthHandler{Service: service}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req struct {
		Email string `json:"email"`
	}

	if err := c.BodyParser(&req); err != nil {
		log.Printf("Invalid request payload: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.CreateResponse(
			"ERROR", "Invalid request payload", nil, nil, nil, nil, nil,
		))
	}

	err := h.Service.GenerateOTP(req.Email)
	if err != nil {
		log.Printf("Failed to generate OTP: %v", err)
		return c.Status(fiber.StatusUnauthorized).JSON(utils.CreateResponse(
			"ERROR", err.Error(), nil, nil, nil, nil, nil,
		))
	}

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
		log.Printf("Invalid request payload: %v", err)
		resp := utils.CreateResponse(
			"ERROR",
			"Invalid request payload",
			nil,
			nil, nil, nil, nil,
		)
		return c.Status(fiber.StatusBadRequest).JSON(resp)
	}

	token, err := h.Service.VerifyOTP(req.Email, req.OTPCode)
	if err != nil {
		log.Printf("Failed to verify OTP: %v", err)
		resp := utils.CreateResponse(
			"ERROR",
			"Invalid or expired OTP",
			nil,
			nil,
			nil,
			nil,
			nil,
		)
		return c.Status(fiber.StatusUnauthorized).JSON(resp)
	}

	resp := utils.CreateResponse(
		"SUCCESS",
		"OTP verified successfully",
		map[string]string{"accessToken": token},
		nil, nil, nil, nil,
	)
	return c.Status(fiber.StatusOK).JSON(resp)
}
