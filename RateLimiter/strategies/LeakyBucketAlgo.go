package algo

type LeakyBucketRateLimiter struct {
	Buffer []int
}

func NewLeakyBucketRateLimiter() IRateLimiter {
	return &LeakyBucketRateLimiter{
		Buffer: make([]int, 0),
	}
}

func (a *LeakyBucketRateLimiter) IsRequestAllowed() bool {
	return true
}
