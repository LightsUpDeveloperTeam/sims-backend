package schoolsmasterdata

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"sims-backend/internal/authentication"
)

func RegisterSchoolRoutes(app *fiber.App, db *gorm.DB) {
	schoolsRepo := NewRepository(db)
	schoolsService := NewService(schoolsRepo)
	schoolsHandler := NewHandler(schoolsService)

	schoolsGroup := app.Group("/schools", authentication.JWTMiddleware())
	schoolsGroup.Post("/create", schoolsHandler.CreateSchools)
	schoolsGroup.Get("/", schoolsHandler.GetSchools)
	schoolsGroup.Get("/:id", schoolsHandler.GetSchoolByID)
	schoolsGroup.Put("/:id", schoolsHandler.UpdateSchool)
	schoolsGroup.Delete("/:id", schoolsHandler.DeleteSchool)
}
