package industry_partner

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"sims-backend/internal/authentication"
)

func IndustryPartnerRoutes(app *fiber.App, db *gorm.DB) {
	industryPartnerRepo := NewRepository(db)
	industryPartnerService := NewService(industryPartnerRepo)
	industryPartnerHandler := NewHandler(industryPartnerService)

	industryPartnerGroup := app.Group("/industry_partner", authentication.JWTMiddleware())
	industryPartnerGroup.Post("/create", industryPartnerHandler.CreateIndustryPartner)
	industryPartnerGroup.Get("/", industryPartnerHandler.GetIndustryPartner)
	industryPartnerGroup.Get("/:id", industryPartnerHandler.GetIndustryPartnerByID)
	industryPartnerGroup.Put("/:id", industryPartnerHandler.UpdateIndustryPartner)
	industryPartnerGroup.Delete("/:id", industryPartnerHandler.DeleteIndustryPartner)
}
