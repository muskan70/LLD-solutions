package algo

import "time"

type TokenBucketLog struct {
	CurrentTokens  int
	LastRefillTime time.Time
}

type TokenBucketRateLimiter struct {
	Store map[int]*TokenBucketLog
}

func NewTokenBucketRateLimiter() *TokenBucketRateLimiter {
	return &TokenBucketRateLimiter{
		Store: make(map[int]*TokenBucketLog),
	}
}

func (a *TokenBucketRateLimiter) IsRequestAllowed() bool {
	return true
}
