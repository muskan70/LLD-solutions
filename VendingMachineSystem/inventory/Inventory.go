package inventory

import "fmt"

type Inventory struct {
	ItemMap map[string]int
}

func NewInventory() *Inventory {
	return &Inventory{ItemMap: make(map[string]int)}
}

func (i *Inventory) AddNewItem(it *Item, qty int) {
	i.ItemMap[it.Name] = qty
}

func (i *Inventory) UpdateInventory(it *Item) {
	i.ItemMap[it.Name] = i.ItemMap[it.Name] - 1
}

func (i *Inventory) CheckItemAvailability(it *Item) bool {
	fmt.Println(i.ItemMap[it.Name])
	return i.ItemMap[it.Name] > 0
}

func (i *Inventory) FillInventory() {
	for itmNm := range ItemMap {
		it := NewItem(itmNm)
		i.ItemMap[it.Name] = 5
	}
}

func (i *Inventory) DisplayInventory() {
	for itm, qty := range i.ItemMap {
		fmt.Println(itm, qty)
	}
}
