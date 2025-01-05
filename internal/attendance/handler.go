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

func (h *AttendanceHandler) ClockIn(c *fiber.Ctx) error {
	var input struct {
		UserID    uint    `json:"user_id"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	record, err := h.Service.ClockIn(input.UserID, input.Latitude, input.Longitude)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(record)
}

func (h *AttendanceHandler) ClockOut(c *fiber.Ctx) error {
	var input struct {
		UserID    uint    `json:"user_id"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	record, err := h.Service.ClockOut(input.UserID, input.Latitude, input.Longitude)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(record)
}
