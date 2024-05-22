package ratelimiter

type RateLimiter struct {
	strategy Strategy
}

func NewRateLimiter(strategy Strategy) *RateLimiter {
	return &RateLimiter{strategy: strategy}
}

func (r *RateLimiter) IsLimitExceeded(ip string, token string) bool {
	if token != "" {
		return r.strategy.IsTokenLimitExceeded(token)
	}
	return r.strategy.IsIPLimitExceeded(ip)
}
