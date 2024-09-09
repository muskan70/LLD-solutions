package inventory

type Inventory struct {
	ItemMap map[*Item]int
}

func NewInventory() *Inventory {
	return &Inventory{ItemMap: make(map[*Item]int)}
}

func (i *Inventory) AddNewItem(it *Item, qty int) {
	i.ItemMap[it] = qty
}

func (i *Inventory) UpdateInventory(it *Item) {
	i.ItemMap[it] = i.ItemMap[it] - 1
}

func (i *Inventory) CheckItemAvailability(it *Item) bool {
	return i.ItemMap[it] > 0
}

func (i *Inventory) FillInventory() {
	cd := 0
	for itmNm, itmDetails := range ItemMap {
		it := NewItem(itmNm, itmDetails[0], cd, itmDetails[1])
		i.ItemMap[it] = 5
		cd++
	}
}
