package inventory

const (
	ITEM_TYPE_CARBONATED_DRINK = iota + 1
	ITEM_TYPE_JUICE
	ITEM_TYPE_COFFEE
	ITEM_TYPE_TEA
	ITEM_TYPE_SNACKS
)

var ItemMap = map[string][]int{
	"Coke":             {1, ITEM_TYPE_CARBONATED_DRINK, 60},
	"Pepsi":            {2, ITEM_TYPE_CARBONATED_DRINK, 55},
	"Apple Juice":      {3, ITEM_TYPE_JUICE, 30},
	"Guava Juice":      {4, ITEM_TYPE_JUICE, 40},
	"Iced Latte":       {5, ITEM_TYPE_COFFEE, 50},
	"Amul Cold Coffee": {6, ITEM_TYPE_COFFEE, 40},
	"Ice Tea":          {7, ITEM_TYPE_TEA, 20},
	"Hot Black Tea":    {8, ITEM_TYPE_TEA, 15},
	"Lays":             {9, ITEM_TYPE_SNACKS, 20},
	"Nachos":           {10, ITEM_TYPE_SNACKS, 20},
	"Bingo":            {11, ITEM_TYPE_SNACKS, 10},
}

type Item struct {
	Name     string
	ItemType int
	ItemCode int
	Price    int
}

func NewItem(nm string) *Item {
	return &Item{
		Name:     nm,
		ItemType: ItemMap[nm][1],
		ItemCode: ItemMap[nm][0],
		Price:    ItemMap[nm][2],
	}
}
