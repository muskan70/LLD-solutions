package main

import (
	algo "rateLimiter/RateLimitingAlgo"
	"rateLimiter/config"
)

type RequestStore map[int]map[int]algo.IRateLimiter

var store RequestStore

func NewStore() {
	store = make(RequestStore)
}

func AddNewRequestStore(apiId, userId int) {
	config := config.GetThrottleConfig(apiId)
	rateLimiter := algo.RateLimiterFactory(config.RateLimiterAlgo)
	store[apiId][userId] = rateLimiter
}
