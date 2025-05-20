package main

// START OMIT
orderEntities := []OrderEntity{ ... }

// Map to protobuf structure.	
orderProtos := SliceMap(orderEntities, func(e OrderEntity) OrderProto {
	return OrderProto{ ... }
})

// Map back to entity object.	
orderEntities = SliceMap(orderProtos, func(e OrderProto) OrderEntity {
	return OrderEntity{ ... }
})
// END OMIT
