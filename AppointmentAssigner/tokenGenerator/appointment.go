package tokenGenerator

import (
	"sync/atomic"
)

var tokenId atomic.Uint64

type Token struct {
	TokenId        uint64
	CustomerId     uint64 // unique identifier
	DepartmentType int
	Status         int // 0- new, 1- completed, 2- in queue
}

var tokenMap map[uint64]*Token

func NewToken(customerId uint64) *Token {
	token := &Token{
		TokenId:    tokenId.Add(1),
		CustomerId: customerId,
		Status:     0,
	}
	tokenMap[token.TokenId] = token
	return token
}

func GetTokenDetails(tokenId uint64) *Token {
	return tokenMap[tokenId]
}

func (c *Token) MarkComplete() {
	c.Status = 1
}
