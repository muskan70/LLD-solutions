package coin

const (
	COIN_TYPE_PENNY   = 1
	COIN_TYPE_NICKEL  = 5
	COIN_TYPE_DIME    = 10
	COIN_TYPE_QUARTER = 25
)

type Coin struct {
	CoinType int
}

func (c *Coin) GetValue() int {
	return c.CoinType
}

func GetTotalCoinsValue(c []*Coin) int {
	total := 0
	for i := range c {
		total = total + c[i].CoinType
	}
	return total
}
