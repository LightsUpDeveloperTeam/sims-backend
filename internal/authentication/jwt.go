package authentication

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var (
	redisClient *redis.Client
	ctx         = context.Background()
)

// initialize Redis
func init() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	if redisHost == "" || redisPort == "" {
		panic("Redis configuration missing in environment")
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
}

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    "ERROR",
				"message": "Authorization header missing",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    "ERROR",
				"message": "Invalid token format",
			})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if isTokenBlacklisted(tokenString) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    "ERROR",
				"message": "Token is blacklisted",
			})
		}

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    "ERROR",
				"message": "Invalid or expired token",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Locals("email", claims["email"])
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"code":    "ERROR",
				"message": "Invalid token claims",
			})
		}
		return c.Next()
	}
}

func isTokenBlacklisted(token string) bool {
	_, err := redisClient.Get(ctx, token).Result()
	return err == nil // Jika token ditemukan, berarti di-blacklist
}