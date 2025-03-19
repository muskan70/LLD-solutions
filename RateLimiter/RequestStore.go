package main

import algo "rateLimiter/RateLimitingAlgo"

type RequestStore struct {
	RateLimiter algo.IRateLimiter
	UserId      int
}

type RequestStoreService map[int]RequestStore
