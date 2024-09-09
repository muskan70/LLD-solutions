package machine

import (
	"errors"
	"vending/coin"
	"vending/inventory"
)

type HasMoneyState struct {
	mchn *VendingMachine
}

func (idl *HasMoneyState) ClickOnInsertCoinButton() error {
	return errors.New("Please insert coins")
}
func (idl *HasMoneyState) InsertCoins(c *coin.Coin) error {
	idl.mchn.AddCoin(c)
	return nil
}
func (idl *HasMoneyState) ClickOnProductSelectionButton() error {
	idl.mchn.SetStatus(&ProductSelectionState{mchn: idl.mchn})
	return nil
}
func (idl *HasMoneyState) ChooseProduct(itm *inventory.Item) error {
	return errors.New("Press Product Selection button to choose product")
}
func (idl *HasMoneyState) CancelButtonOrRefundCoins() error {
	idl.mchn.RefundCoins()
	return nil
}
func (idl *HasMoneyState) DispenseProduct(itm *inventory.Item) error {
	return errors.New("Product cannot be dispensed at this state")
}
