package machine

import (
	"errors"
	"fmt"
	"sync"
	"vending/coin"
	"vending/inventory"
)

type VendingMachine struct {
	invnt  *inventory.Inventory
	coins  []*coin.Coin
	status IMachineState
}

var vdMch *VendingMachine
var lock = sync.Mutex{}

func GetVendingMachineInstance() *VendingMachine {
	if vdMch == nil {
		lock.Lock()
		defer lock.Unlock()
		if vdMch == nil {
			vdMch = &VendingMachine{
				invnt:  inventory.NewInventory(),
				coins:  []*coin.Coin{},
				status: &IdleState{},
			}
		}
	}
	return vdMch
}

func (v *VendingMachine) SetStatus(s IMachineState) {
	v.status = s
}

func (v *VendingMachine) UpdateInventory(itm *inventory.Item) {
	v.invnt.UpdateInventory(itm)
	v.SetStatus(&IdleState{})
}

func (v *VendingMachine) AddCoin(c *coin.Coin) {
	v.coins = append(v.coins, c)
}

func (v *VendingMachine) RefundCoins() {
	fmt.Println("All inserted coins have been refunded")
	v.coins = []*coin.Coin{}
	v.SetStatus(&IdleState{})
}

func (v *VendingMachine) CheckProductAvailabiltyNPrice(itm *inventory.Item) error {
	if !v.invnt.CheckItemAvailability(itm) {
		return errors.New("Item not Available")
	}
	if coin.GetTotalCoinsValue(v.coins) < itm.Price {
		return errors.New("Insufficient Coins")
	}
	return nil
}
