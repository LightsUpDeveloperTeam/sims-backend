package attendance

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"sims-backend/internal/authentication"
)

func RegisterAttendanceRoutes(app *fiber.App, db *gorm.DB) {
	attendanceRepo := NewAttendanceRepository(db)
	attendanceService := NewAttendanceService(attendanceRepo)
	attendanceHandler := NewAttendanceHandler(attendanceService)

	attendanceGroup := app.Group("/attendance", authentication.JWTMiddleware())
	attendanceGroup.Post("/clock", attendanceHandler.ClockInOut) // Unified clock-in and clock-out route
}
