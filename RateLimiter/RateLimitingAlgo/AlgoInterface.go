package algo

import "rateLimiter/constants"

type IRateLimiter interface {
	IsRequestAllowed()
}

func RateLimiterFactory(algo int) IRateLimiter {
	if algo == constants.RATE_LIMITER_TOKEN_BUCKET {
		return &TokenBucketAlgo{}
	} else if algo == constants.RATE_LIMITER_LEAKY_BUCKET {
		return &LeakyBucketAlgo{}
	} else if algo == constants.RATE_LIMITER_FIXED_WINDOW_COUNTER {
		return &FixedWindowCounterAlgo{}
	} else if algo == constants.RATE_LIMITER_SLIDING_WINDOW_LOG {
		return &SlidingWindowLogAlgo{}
	} else if algo == constants.RATE_LIMITER_SLIDING_WINDOW_COUNTER {
		return &SlidingWindowCounterAlgo{}
	}
	return nil
}
