package machine

import (
	"errors"
	"fmt"
	"vending/coin"
	"vending/inventory"
)

type DispenseProductState struct {
	mchn *VendingMachine
}

func (dp *DispenseProductState) ClickOnInsertCoinButton() error {
	return errors.New("Click on insert coin invalid at dispense product State")
}
func (dp *DispenseProductState) InsertCoins(c *coin.Coin) error {
	return errors.New("Insert coins not allowed at dispense product State")
}
func (dp *DispenseProductState) ClickOnProductSelectionButton() error {
	return errors.New("Already product selection done")
}
func (dp *DispenseProductState) ChooseProduct(itm *inventory.Item) error {
	return errors.New("Already product selection done")
}
func (dp *DispenseProductState) CancelButtonOrRefundCoins() error {
	return errors.New("Refund coins not allowed at dispense product State")
}
func (dp *DispenseProductState) DispenseProduct(itm *inventory.Item) error {
	fmt.Println("Product dispensed")
	dp.mchn.UpdateInventory(itm)
	return nil
}
