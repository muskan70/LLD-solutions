package algo

import "time"

type SlidingWindowLog struct {
	RequestTimestamps []time.Time
}

type SlidingWindowLogRateLimiter struct {
	Store map[int]*SlidingWindowLog
}

func NewSlidingWindowLogRateLimiter() *SlidingWindowLogRateLimiter {
	return &SlidingWindowLogRateLimiter{
		Store: make(map[int]*SlidingWindowLog),
	}
}

func (a *SlidingWindowLogRateLimiter) IsRequestAllowed() bool {
	return true
}
