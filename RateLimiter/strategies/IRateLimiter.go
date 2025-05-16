package algo

import "rateLimiter/constants"

type IRateLimiter interface {
	IsRequestAllowed() bool
}

func RateLimiterFactory(algo int) IRateLimiter {
	if algo == constants.RATE_LIMITER_TOKEN_BUCKET {
		return NewTokenBucketRateLimiter()
	} else if algo == constants.RATE_LIMITER_LEAKY_BUCKET {
		return NewLeakyBucketRateLimiter()
	} else if algo == constants.RATE_LIMITER_FIXED_WINDOW_COUNTER {
		return NewFixedWindowCounterRateLimiter()
	} else if algo == constants.RATE_LIMITER_SLIDING_WINDOW_LOG {
		return NewSlidingWindowLogRateLimiter()
	} else if algo == constants.RATE_LIMITER_SLIDING_WINDOW_COUNTER {
		return NewSlidingWindowCounterRateLimiter()
	}
	return nil
}
