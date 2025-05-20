package main

type Order struct {
	OrderID int64
	Name    string
}

type OrderItem struct {
	ItemID int64
	Name   string
}

// START OMIT
func GetOrderIDFromOrders(orders []Order) []int64 {
	orderIDs := make([]int64, len(orders))
	for i, order := range orders {
		orderIDs[i] = order.OrderID
	}
	return orderIDs
}

func GetItemIDFromOrderItems(orderItems []OrderItem) []int64 {
	itemIDs := make([]int64, len(orderItems))
	for i, item := range orderItems {
		itemIDs[i] = item.ItemID
	}
	return itemIDs
}

// END OMIT
