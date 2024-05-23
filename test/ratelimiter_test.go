package test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/rzeradev/rate-limiter/configs"
	"github.com/rzeradev/rate-limiter/internal/middleware"
	"github.com/rzeradev/rate-limiter/pkg/ratelimiter"
	"github.com/stretchr/testify/assert"
)

func TestRateLimiter(t *testing.T) {
	envPath := "/../.env"
	configs.LoadConfig(&envPath)

	app := fiber.New()

	redisStrategy := ratelimiter.NewRedisStrategy()
	limiter := ratelimiter.NewRateLimiter(redisStrategy)
	rateLimiterMiddleware := middleware.NewRateLimiterMiddleware(limiter)

	app.Use(rateLimiterMiddleware.Limit())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the rate limited API!")
	})

	t.Run("Test IP based rate limiting", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)

		for i := 0; i < 10; i++ {
			resp, _ := app.Test(req)
			if i < 5 {
				assert.Equal(t, 200, resp.StatusCode)
			} else {
				assert.Equal(t, 429, resp.StatusCode)
			}
		}
	})

	t.Run("Test API_KEY based rate limiting", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("API_KEY", "test-token")

		for i := 0; i < 10; i++ {
			resp, _ := app.Test(req)
			if i < 10 {
				assert.Equal(t, 200, resp.StatusCode)
			} else {
				assert.Equal(t, 429, resp.StatusCode)
			}
		}
	})
}
