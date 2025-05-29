package tokenGenerator

import (
	"fmt"
	"sync/atomic"
)

var tokenId atomic.Uint64

type Token struct {
	TokenId        uint64
	CustomerId     int // unique identifier
	DepartmentType int
	Status         int // 0- new, 1- completed, 2- in queue
}

type TokenManager struct {
	TokenMap map[uint64]*Token
}

func NewTokenManager() *TokenManager {
	tm := &TokenManager{TokenMap: make(map[uint64]*Token)}
	return tm
}

func (tm *TokenManager) NewToken(customerId int) *Token {
	token := &Token{
		TokenId:    tokenId.Add(1),
		CustomerId: customerId,
		Status:     0,
	}
	tm.TokenMap[token.TokenId] = token
	return token
}

func (tm *TokenManager) GetTokenDetails(tokenId uint64) *Token {
	return tm.TokenMap[tokenId]
}

func (c *Token) MarkComplete() {
	fmt.Println("task completed", c.TokenId)
	c.Status = 1
}
