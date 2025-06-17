package entity

type Book struct {
	BookID   int     `json:"id"`
	Name     string  `json:"name"`
	Writer   string  `json:"author"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"` //tells about total books initially
	Sold     int     `json:"sold"`
	InStock  int     `json:"in_stock"`
	Status   string  `json:"status"`
}
