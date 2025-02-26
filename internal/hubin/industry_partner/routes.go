package industry_partner

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"sims-backend/internal/authentication"
)

func CollaborationHistoryRoutes(app *fiber.App, db *gorm.DB) {
	collaborationHistoryRepo := NewRepository(db)
	collaborationHistoryService := NewService(collaborationHistoryRepo)
	collaborationHistoryHandler := NewHandler(collaborationHistoryService)

	collaborationHistoryGroup := app.Group("/collaboration_history", authentication.JWTMiddleware())
	collaborationHistoryGroup.Post("/create", collaborationHistoryHandler.CreateCollaborationHistory)
	collaborationHistoryGroup.Get("/", collaborationHistoryHandler.GetCollaborationHistory)
	collaborationHistoryGroup.Get("/:id", collaborationHistoryHandler.GetCollaborationHistoryByID)
	collaborationHistoryGroup.Put("/:id", collaborationHistoryHandler.UpdateCollaborationHistory)
	collaborationHistoryGroup.Delete("/:id", collaborationHistoryHandler.DeleteCollaborationHistory)
}

func MemorandumOfUnderstandingRoutes(app *fiber.App, db *gorm.DB) {
	memorandumOfUnderstandingRepo := NewRepository(db)
	memorandumOfUnderstandingService := NewService(memorandumOfUnderstandingRepo)
	memorandumOfUnderstandingHandler := NewHandler(memorandumOfUnderstandingService)

	memorandumOfUnderstandingGroup := app.Group("/memorandum_of_understanding", authentication.JWTMiddleware())
	memorandumOfUnderstandingGroup.Post("/create", memorandumOfUnderstandingHandler.CreateMemorandumOfUnderstanding)
	memorandumOfUnderstandingGroup.Get("/", memorandumOfUnderstandingHandler.GetMemorandumOfUnderstanding)
	memorandumOfUnderstandingGroup.Get("/:id", memorandumOfUnderstandingHandler.GetMemorandumOfUnderstandingByID)
	memorandumOfUnderstandingGroup.Put("/:id", memorandumOfUnderstandingHandler.UpdateMemorandumOfUnderstanding)
	memorandumOfUnderstandingGroup.Delete("/:id", memorandumOfUnderstandingHandler.DeleteMemorandumOfUnderstanding)
}
