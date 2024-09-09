package machine

import (
	"vending/coin"
	"vending/inventory"
)

type IMachineState interface {
	ClickOnInsertCoinButton() error
	InsertCoins(c *coin.Coin) error
	ClickOnProductSelectionButton() error
	ChooseProduct(itm *inventory.Item) error
	CancelButtonOrRefundCoins() error
	DispenseProduct(itm *inventory.Item) error
}
