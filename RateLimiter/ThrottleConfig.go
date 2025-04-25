package main

type ThrottleConfig struct {
	RateLimiterAlgo     int
	NoofRequestsAllowed int
	WindowInSeconds     int
	MaxBucketCapacity   int //optional
}

var throttleConfigMap map[int]*ThrottleConfig

func NewThrottleConfigMap() {
	throttleConfigMap = make(map[int]*ThrottleConfig)
}

func AddThrottleConfig(apiID, algo, requests, window int) {
	throttleConfigMap[apiID] = &ThrottleConfig{
		RateLimiterAlgo:     algo,
		NoofRequestsAllowed: requests,
		WindowInSeconds:     window,
	}
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

func UpdateRateLimitingAlgo(apiID, algo int) {
	throttleConfigMap[apiID].RateLimiterAlgo = algo
}
