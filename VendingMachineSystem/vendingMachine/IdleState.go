package machine

import (
	"errors"
	"vending/coin"
	"vending/inventory"
)

type IdleState struct {
	mchne *VendingMachine
}

func (idl *IdleState) ClickOnInsertCoinButton() error {
	idl.mchne.SetStatus(&HasMoneyState{mchn: idl.mchne})
	return nil
}
func (idl *IdleState) InsertCoins(c *coin.Coin) error {
	return errors.New("First you need to press insert coin button")
}

func (idl *IdleState) ClickOnProductSelectionButton() error {
	return errors.New("First you need to press insert coin button")
}
func (idl *IdleState) ChooseProduct(itm *inventory.Item) error {
	return errors.New("You cannot choose product in idle state")
}
func (idl *IdleState) CancelButtonOrRefundCoins() error {
	return errors.New("No refund possible in idle state")
}
func (idl *IdleState) DispenseProduct(itm *inventory.Item) error {
	return errors.New("No product dispense possible in idle state")
}
