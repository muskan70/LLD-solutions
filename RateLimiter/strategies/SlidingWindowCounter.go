package algo

import "time"

type SlidingWindowCounterLog struct {
	RequestCounter          int
	WindowStartingTimestamp time.Time
}

type SlidingWindowCounterRateLimiter struct {
	Store map[int]*SlidingWindowCounterLog
}

func NewSlidingWindowCounterRateLimiter() *SlidingWindowCounterRateLimiter {
	return &SlidingWindowCounterRateLimiter{
		Store: make(map[int]*SlidingWindowCounterLog),
	}
}

func (a *SlidingWindowCounterRateLimiter) IsRequestAllowed() bool {
	return true
}
