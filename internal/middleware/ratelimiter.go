package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rzeradev/rate-limiter/pkg/ratelimiter"
)

type RateLimiterMiddleware struct {
	limiter *ratelimiter.RateLimiter
}

func NewRateLimiterMiddleware(limiter *ratelimiter.RateLimiter) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{limiter: limiter}
}

func (m *RateLimiterMiddleware) Limit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		token := c.Get("API_KEY")

		if m.limiter.IsLimitExceeded(ip, token) {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "you have reached the maximum number of requests or actions allowed within a certain time frame",
			})
		}

		return c.Next()
	}
}
