package main

// START OMIT
orderIDs := SliceMap(orders, func(order Order) int64 { return order.OrderID })

orderItemIDs := SliceMap(orderItems, func(item OrderItem) int64 { return item.ItemID })
// END OMIT
