package internship_information

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"sims-backend/internal/authentication"
)

//Internship Contract Routes

func InternshipContractRoutes(app *fiber.App, db *gorm.DB) {
	internshipContractRepo := NewRepository(db)
	internshipContractService := NewService(internshipContractRepo)
	internshipContractHandler := NewHandler(internshipContractService)

	internshipContractGroup := app.Group("/internship_contract", authentication.JWTMiddleware())
	internshipContractGroup.Post("/create", internshipContractHandler.CreateInternshipContract)
	internshipContractGroup.Get("/", internshipContractHandler.GetInternshipContract)
	internshipContractGroup.Get("/:id", internshipContractHandler.GetInternshipContractByID)
	internshipContractGroup.Put("/:id", internshipContractHandler.UpdateInternshipContract)
	internshipContractGroup.Delete("/:id", internshipContractHandler.DeleteInternshipContract)
}

//Internship Document Routes

func InternshipDocumentRoutes(app *fiber.App, db *gorm.DB) {
	internshipDocumentRepo := NewRepository(db)
	internshipDocumentService := NewService(internshipDocumentRepo)
	internshipDocumentHandler := NewHandler(internshipDocumentService)

	internshipDocumentGroup := app.Group("/internship_document", authentication.JWTMiddleware())
	internshipDocumentGroup.Post("/create", internshipDocumentHandler.CreateInternshipDocument)
	internshipDocumentGroup.Get("/", internshipDocumentHandler.GetInternshipDocument)
	internshipDocumentGroup.Get("/:id", internshipDocumentHandler.GetInternshipDocumentByID)
	internshipDocumentGroup.Put("/:id", internshipDocumentHandler.UpdateInternshipDocument)
	internshipDocumentGroup.Delete("/:id", internshipDocumentHandler.DeleteInternshipDocument)
}

//Internship Evaluation Routes

func InternshipEvaluationRoutes(app *fiber.App, db *gorm.DB) {
	internshipEvaluationRepo := NewRepository(db)
	internshipEvaluationService := NewService(internshipEvaluationRepo)
	internshipEvaluationHandler := NewHandler(internshipEvaluationService)

	internshipEvaluationGroup := app.Group("/internship_evaluation", authentication.JWTMiddleware())
	internshipEvaluationGroup.Post("/create", internshipEvaluationHandler.CreateInternshipEvaluation)
	internshipEvaluationGroup.Get("/", internshipEvaluationHandler.GetInternshipEvaluation)
	internshipEvaluationGroup.Get("/:id", internshipEvaluationHandler.GetInternshipEvaluationByID)
	internshipEvaluationGroup.Put("/:id", internshipEvaluationHandler.UpdateInternshipEvaluation)
	internshipEvaluationGroup.Delete("/:id", internshipEvaluationHandler.DeleteInternshipEvaluation)
}

//Internship Progress Routes

func InternshipProgressRoutes(app *fiber.App, db *gorm.DB) {
	internshipProgressRepo := NewRepository(db)
	internshipProgressService := NewService(internshipProgressRepo)
	internshipProgressHandler := NewHandler(internshipProgressService)

	internshipProgressGroup := app.Group("/internship_progress", authentication.JWTMiddleware())
	internshipProgressGroup.Post("/create", internshipProgressHandler.CreateInternshipProgress)
	internshipProgressGroup.Get("/", internshipProgressHandler.GetInternshipProgress)
	internshipProgressGroup.Get("/:id", internshipProgressHandler.GetInternshipProgressByID)
	internshipProgressGroup.Put("/:id", internshipProgressHandler.UpdateInternshipProgress)
	internshipProgressGroup.Delete("/:id", internshipProgressHandler.DeleteInternshipProgress)
}
