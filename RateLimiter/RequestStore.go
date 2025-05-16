package main

import (
	"rateLimiter/config"
	algo "rateLimiter/strategies"
)

var store map[int]algo.IRateLimiter

func NewStore() {
	store = make(map[int]algo.IRateLimiter)
}

func AddRequestStoreForAPI(apiId int) {
	config := config.GetThrottleConfig(apiId)
	rateLimiter := algo.RateLimiterFactory(config.RateLimiterAlgo)
	store[apiId] = rateLimiter
}
