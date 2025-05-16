package algo

type FixedWindowCounterLog struct {
	RequestCounter int
	Window         int
}

type FixedWindowCounterRateLimiter struct {
	Store map[int]*FixedWindowCounterLog
}

func NewFixedWindowCounterRateLimiter() *FixedWindowCounterRateLimiter {
	return &FixedWindowCounterRateLimiter{
		Store: make(map[int]*FixedWindowCounterLog),
	}
}

func (a *FixedWindowCounterRateLimiter) IsRequestAllowed() bool {
	return true
}
