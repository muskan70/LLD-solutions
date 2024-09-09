package machine

import (
	"errors"
	"vending/coin"
	"vending/inventory"
)

type ProductSelectionState struct {
	mchn *VendingMachine
}

func (ps *ProductSelectionState) ClickOnInsertCoinButton() error {
	return errors.New("Click on insert coin invalid at product selection State")
}
func (ps *ProductSelectionState) InsertCoins(c *coin.Coin) error {
	return errors.New("Insert coins not allowed at product selection State")
}
func (ps *ProductSelectionState) ClickOnProductSelectionButton() error {
	return errors.New("Please select a product")
}
func (ps *ProductSelectionState) ChooseProduct(itm *inventory.Item) error {
	err := ps.mchn.CheckProductAvailabiltyNPrice(itm)
	if err != nil {
		return err
	}
	ps.mchn.SetStatus(&DispenseProductState{})
	return nil
}
func (idl *ProductSelectionState) CancelButtonOrRefundCoins() error {
	idl.mchn.RefundCoins()
	return nil
}
func (idl *ProductSelectionState) DispenseProduct(itm *inventory.Item) error {
	return errors.New("Product cannot be dispensed at this state")
}
