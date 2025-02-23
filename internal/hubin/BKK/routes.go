package BKK

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"sims-backend/internal/authentication"
)

func InternshipVacancyRoutes(app *fiber.App, db *gorm.DB) {
	InternshipVacancyRepo := NewRepository(db)
	InternshipVacancyService := NewService(InternshipVacancyRepo)
	InternshipVacancyHandler := NewHandler(InternshipVacancyService)

	InternshipVacancyGroup := app.Group("/vacancy", authentication.JWTMiddleware())
	InternshipVacancyGroup.Post("/create", InternshipVacancyHandler.CreateVacancy)
	InternshipVacancyGroup.Get("/", InternshipVacancyHandler.GetVacancies)
	InternshipVacancyGroup.Get("/:id", InternshipVacancyHandler.GetVacancyByID)
	InternshipVacancyGroup.Put("/:id", InternshipVacancyHandler.UpdateVacancy)
	InternshipVacancyGroup.Delete("/:id", InternshipVacancyHandler.DeleteVacancy)
}

func InternshipRegistrationRoutes(app *fiber.App, db *gorm.DB) {
	InternshipRegistrationRepo := NewRepository(db)
	InternshipRegistrationService := NewService(InternshipRegistrationRepo)
	InternshipRegistrationHandler := NewHandler(InternshipRegistrationService)

	InternshipRegistrationGroup := app.Group("/vacancy", authentication.JWTMiddleware())
	InternshipRegistrationGroup.Post("/create", InternshipRegistrationHandler.CreateRegistration)
	InternshipRegistrationGroup.Get("/", InternshipRegistrationHandler.GetVacancies)
	InternshipRegistrationGroup.Get("/:id", InternshipRegistrationHandler.GetRegistrationByID)
	InternshipRegistrationGroup.Delete("/:id", InternshipRegistrationHandler.DeleteRegistration)
}

func AlumnusDistributionRoutes(app *fiber.App, db *gorm.DB) {
	AlumnusDistributionRepo := NewRepository(db)
	AlumnusDistributionService := NewService(AlumnusDistributionRepo)
	AlumnusDistributionHandler := NewHandler(AlumnusDistributionService)

	AlumnusDistributionGroup := app.Group("/vacancy", authentication.JWTMiddleware())
	AlumnusDistributionGroup.Post("/create", AlumnusDistributionHandler.CreateDistribution)
	AlumnusDistributionGroup.Get("/", AlumnusDistributionHandler.GetVacancies)
	AlumnusDistributionGroup.Get("/:id", AlumnusDistributionHandler.GetDistributionByID)
	AlumnusDistributionGroup.Put("/:id", AlumnusDistributionHandler.UpdateDistribution)
	AlumnusDistributionGroup.Delete("/:id", AlumnusDistributionHandler.DeleteDistribution)
}
