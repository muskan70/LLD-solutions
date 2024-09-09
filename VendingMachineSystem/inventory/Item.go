package inventory

const (
	ITEM_TYPE_CARBONATED_DRINK = iota + 1
	ITEM_TYPE_JUICE
	ITEM_TYPE_COFFEE
	ITEM_TYPE_TEA
	ITEM_TYPE_SNACKS
)

var ItemMap = map[string][]int{
	"Coke":             {ITEM_TYPE_CARBONATED_DRINK, 60},
	"Pepsi":            {ITEM_TYPE_CARBONATED_DRINK, 55},
	"Apple Juice":      {ITEM_TYPE_JUICE, 30},
	"Guava Juice":      {ITEM_TYPE_JUICE, 40},
	"Iced Latte":       {ITEM_TYPE_COFFEE, 50},
	"Amul Cold Coffee": {ITEM_TYPE_COFFEE, 40},
	"Ice Tea":          {ITEM_TYPE_TEA, 20},
	"Hot Black Tea":    {ITEM_TYPE_TEA, 15},
	"Lays":             {ITEM_TYPE_SNACKS, 20},
	"Nachos":           {ITEM_TYPE_SNACKS, 20},
	"Bingo":            {ITEM_TYPE_SNACKS, 10},
}

type Item struct {
	Name     string
	ItemType int
	ItemCode int
	Price    int
}

func NewItem(nm string, it int, c int, p int) *Item {
	return &Item{
		Name:     nm,
		ItemType: it,
		ItemCode: c,
		Price:    p,
	}
}
