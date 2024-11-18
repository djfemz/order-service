package models


type Order struct{
	Id uint64
	Item string
	Price float64
	CreatedAt string
	Status string
}

func NewOrder(item string, price float64) *Order {
	return &Order{
		Item: item,
		Price: price,
	}
}