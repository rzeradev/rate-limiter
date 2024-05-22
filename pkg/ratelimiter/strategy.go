package ratelimiter

type Strategy interface {
	IsIPLimitExceeded(ip string) bool
	IsTokenLimitExceeded(token string) bool
}
