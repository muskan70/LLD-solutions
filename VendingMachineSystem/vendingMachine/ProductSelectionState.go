package machine

import (
	"errors"
	"fmt"
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
		fmt.Println(err)
		ps.mchn.SetStatus(&IdleState{mchne: ps.mchn})
		ps.mchn.RefundCoins()
		return err
	}
	fmt.Println("Product selected")
	ps.mchn.SetStatus(&DispenseProductState{mchn: ps.mchn})
	coins := ps.mchn.ReturnChange(itm)
	if coins > 0 {
		fmt.Println("Returned remaining change:", coins)
	}
	return nil
}
func (ps *ProductSelectionState) CancelButtonOrRefundCoins() error {
	ps.mchn.RefundCoins()
	return nil
}
func (ps *ProductSelectionState) DispenseProduct(itm *inventory.Item) error {
	return errors.New("Product cannot be dispensed at this state")
}
