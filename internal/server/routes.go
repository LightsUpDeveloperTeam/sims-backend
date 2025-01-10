package server

import (
	"log"
	"sims-backend/internal/attendance"
	"sims-backend/internal/authentication"
	schoolsmasterdata "sims-backend/internal/masters/schools"
	usersmasterdata "sims-backend/internal/masters/users"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

// RegisterFiberRoutes registers all routes for the application
func RegisterFiberRoutes(app *fiber.App, db *gorm.DB) {
	// Middleware for CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type",
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Public routes
	app.Get("/", HelloWorldHandler)
	app.Get("/health", healthHandler)
	app.Get("/websocket", websocket.New(websocketHandler))

	// Authentication routes
	authHandler := authentication.NewAuthHandler(db)
	app.Post("/auth/login", authHandler.Login)
	app.Post("/auth/verify-otp", authHandler.VerifyOTP)
	app.Post("/auth/refresh-token", authentication.JWTMiddleware(), authHandler.RefreshToken)
	app.Post("/auth/logout", authentication.JWTMiddleware(), authHandler.Logout)

	// Modular routes
	attendance.RegisterAttendanceRoutes(app, db)
	schoolsmasterdata.RegisterSchoolRoutes(app, db)
	usersmasterdata.RegisterUserRoutes(app, db)
}

// Handlers for basic routes
func HelloWorldHandler(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"message": "Hello World"})
}

func healthHandler(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"status": "healthy"})
}

func websocketHandler(con *websocket.Conn) {
	for {
		messageType, message, err := con.ReadMessage()
		if err != nil {
			log.Println("WebSocket Error:", err)
			break
		}
		log.Printf("Received message: %s", message)
		con.WriteMessage(messageType, []byte("Echo: "+string(message)))
	}
}
