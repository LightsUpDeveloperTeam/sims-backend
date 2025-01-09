package server

import (
	"sims-backend/internal/database"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	*fiber.App
	db database.Service
}

// New initializes the Fiber server with middlewares
func New() *FiberServer {
	return &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "sims-backend",
			AppName:      "sims-backend",
		}),
		db: database.New(),
	}
}

// RegisterRoutes initializes routes from routes.go
func (s *FiberServer) RegisterFiberRoutes() {
	RegisterFiberRoutes(s.App, s.db.GetGORMDB()) // Fixed: Use GetGORMDB() instead of GetDB()
}
