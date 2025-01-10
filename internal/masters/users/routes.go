package usersmasterdata

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"sims-backend/internal/authentication"
)

func RegisterUserRoutes(app *fiber.App, db *gorm.DB) {
	userRepo := NewRepository(db)
	userService := NewService(userRepo)
	userHandler := NewHandler(userService)

	usersGroup := app.Group("/users", authentication.JWTMiddleware())
	usersGroup.Post("/", userHandler.CreateUser)
	usersGroup.Put("/:id", userHandler.UpdateUser)
	usersGroup.Get("/:id", userHandler.GetUserByID)
	usersGroup.Delete("/:id", userHandler.DeleteUser)

	// Role and permission management
	usersGroup.Post("/roles", userHandler.CreateRole)
	usersGroup.Put("/roles/:id", userHandler.UpdateRole)
	usersGroup.Delete("/roles/:id", userHandler.DeleteRole)
	usersGroup.Post("/permissions", userHandler.CreatePermission)
	usersGroup.Post("/roles/assign-permission", userHandler.AssignPermissionToRole)
}
