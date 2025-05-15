package config

import "rateLimiter/constants"

type ThrottleConfig struct {
	API_Id              int
	RateLimiterAlgo     int
	NoofRequestsAllowed int
	WindowInSeconds     int  // consider refiller rate in token bucket
	MaxBucketCapacity   *int // in token bucket and leaky bucket
}

var throttleConfigMap map[int]*ThrottleConfig

func NewThrottleConfigMap() {
	throttleConfigMap = make(map[int]*ThrottleConfig)
}

func AddThrottleConfig(apiID, algo, requests, window int, bucketSize *int) {
	conf := &ThrottleConfig{
		RateLimiterAlgo:     algo,
		NoofRequestsAllowed: requests,
		WindowInSeconds:     window,
	}
	if algo == constants.RATE_LIMITER_TOKEN_BUCKET || algo == constants.RATE_LIMITER_LEAKY_BUCKET {
		conf.MaxBucketCapacity = bucketSize
	}
	throttleConfigMap[apiID] = conf
}

func GetThrottleConfig(apiID int) *ThrottleConfig {
	config, ok := throttleConfigMap[apiID]
	if !ok {
		return nil
	}
	return config
}

func UpdateAllowedRequests(apiID, requests int) {
	throttleConfigMap[apiID].NoofRequestsAllowed = requests
}

func UpdateWindow(apiID, window int) {
	throttleConfigMap[apiID].WindowInSeconds = window
}

func UpdateRateLimitingAlgo(apiID, algo int, bucketSize *int) {
	throttleConfigMap[apiID].RateLimiterAlgo = algo
	if algo == constants.RATE_LIMITER_TOKEN_BUCKET || algo == constants.RATE_LIMITER_LEAKY_BUCKET {
		throttleConfigMap[apiID].MaxBucketCapacity = bucketSize
	}
}
