package server

import (
	"github.com/gofiber/fiber/v2"

	"sims-backend/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "sims-backend",
			AppName:      "sims-backend",
		}),

		db: database.New(),
	}

	return server
}
