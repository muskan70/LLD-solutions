package algo

import "time"

type TokenBucketAlgo struct {
	CurrentTokens  int
	LastRefillTime time.Time
}

func (a *TokenBucketAlgo) IsRequestAllowed()
