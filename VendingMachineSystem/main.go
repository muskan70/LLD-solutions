package main

import (
	"fmt"
	"vending/coin"
	"vending/inventory"
	machine "vending/vendingMachine"
)

func main() {
	mach := machine.GetVendingMachineInstance()

	mach.FillInventory()
	mach.DisplayInventory()

	state := mach.GetStatus()
	fmt.Println(state.ClickOnInsertCoinButton())

	state = mach.GetStatus()
	fmt.Println(state.InsertCoins(&coin.Coin{CoinType: coin.COIN_TYPE_DIME}))
	fmt.Println(state.InsertCoins(&coin.Coin{CoinType: coin.COIN_TYPE_QUARTER}))
	fmt.Println(state.InsertCoins(&coin.Coin{CoinType: coin.COIN_TYPE_QUARTER}))
	fmt.Println(state.InsertCoins(&coin.Coin{CoinType: coin.COIN_TYPE_NICKEL}))
	fmt.Println(state.ClickOnProductSelectionButton())

	itm := inventory.NewItem("Coke")
	state = mach.GetStatus()
	fmt.Println(state.ChooseProduct(itm))

	state = mach.GetStatus()
	fmt.Println(state.DispenseProduct(itm))
}
