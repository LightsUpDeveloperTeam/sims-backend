package attendance

import (
	"github.com/gofiber/fiber/v2"
)

type AttendanceHandler struct {
	Service *AttendanceService
}

func NewAttendanceHandler(service *AttendanceService) *AttendanceHandler {
	return &AttendanceHandler{Service: service}
}

func (h *AttendanceHandler) ClockInOut(c *fiber.Ctx) error {
	var input struct {
		UserID              uint    `json:"user_id"`
		Latitude            float64 `json:"latitude"`
		Longitude           float64 `json:"longitude"`
		EarlyClockOutReason string  `json:"early_clockout_reason,omitempty"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Panggil service untuk menangani logika clock-in atau clock-out
	record, err := h.Service.ClockInOut(input.UserID, input.Latitude, input.Longitude, input.EarlyClockOutReason)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(record)
}
