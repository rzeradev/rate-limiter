package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rzeradev/rate-limiter/configs"
	"github.com/rzeradev/rate-limiter/internal/middleware"
	"github.com/rzeradev/rate-limiter/pkg/ratelimiter"
)

func main() {
	// Load configuration
	configs.LoadConfig(nil)

	// Create a new Fiber app
	app := fiber.New()

	// Initialize rate limiter with Redis strategy
	redisStrategy := ratelimiter.NewRedisStrategy()
	limiter := ratelimiter.NewRateLimiter(redisStrategy)
	middleware := middleware.NewRateLimiterMiddleware(limiter)

	// Apply middleware
	app.Use(middleware.Limit())

	// Define a test route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the rate limited API!")
	})

	// Start the server
	app.Listen(":8080")
}
