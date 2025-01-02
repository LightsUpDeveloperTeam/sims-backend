package server

import (
	"context"
	"fmt"
	"log"
	"sims-backend/internal/authentication"
	schools "sims-backend/internal/schools-master-data"
	"sims-backend/internal/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/contrib/websocket"
)

func (s *FiberServer) RegisterFiberRoutes() {
	// Apply CORS middleware
	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type",
		AllowCredentials: false, // credentials require explicit origins
		MaxAge:           300,
	}))

	s.App.Get("/", s.HelloWorldHandler)

	s.App.Get("/health", s.healthHandler)

	s.App.Get("/websocket", websocket.New(s.websocketHandler))

	authHandler := authentication.NewAuthHandler(s.db.GetGORMDB()) 
	s.App.Post("/auth/login", authHandler.Login)                   
	s.App.Post("/auth/verify-otp", authHandler.VerifyOTP)            
	s.App.Post("/auth/refresh-token", authentication.JWTMiddleware(), authHandler.RefreshToken) 
	s.App.Post("/auth/logout", authentication.JWTMiddleware(), authHandler.Logout)              	

	protected := s.App.Group("/protected", authentication.JWTMiddleware())
	protected.Get("/profile", func(c *fiber.Ctx) error {
		email := c.Locals("email")
		resp := utils.CreateResponse(
			"SUCCESS",
			"Welcome to your account",
			map[string]interface{}{"email": email},
			nil, // No error code	
			nil, // No error message
			nil, // No error details
			nil, // No pagination
		)
		
		return c.JSON(resp)
	})


	schoolsRepo := schools.NewRepository(s.db.GetGORMDB())
	schoolsService := schools.NewService(schoolsRepo)
	schoolsHandler := schools.NewHandler(schoolsService)

	schools := s.App.Group("/schools", authentication.JWTMiddleware())
	schools.Post("/create", schoolsHandler.CreateSchools)      
	schools.Get("/", schoolsHandler.GetSchools)       
	schools.Get("/:id", schoolsHandler.GetSchoolByID)   
	schools.Put("/:id", schoolsHandler.UpdateSchool)    
	schools.Delete("/:id", schoolsHandler.DeleteSchool) 
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	// Use the global response function
	resp := utils.CreateResponse(
		"SUCCESS",
		"Hello World retrieved successfully.",
		map[string]interface{}{
			"message": "Hello World",
		},
		nil, // No error code
		nil, // No error message
		nil, // No error details
		nil, // No pagination
	)

	return c.JSON(resp)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}

func (s *FiberServer) websocketHandler(con *websocket.Conn) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			_, _, err := con.ReadMessage()
			if err != nil {
				cancel()
				log.Println("Receiver Closing", err)
				break
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
			if err := con.WriteMessage(websocket.TextMessage, []byte(payload)); err != nil {
				log.Printf("could not write to socket: %v", err)
				return
			}
			time.Sleep(time.Second * 2)
		}
	}
}
