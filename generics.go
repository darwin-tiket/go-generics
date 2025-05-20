package main

// START OMIT
type Integer interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

func SumValues[T Integer](iter []T) T {
	var sum T
	for _, v := range iter {
		sum += v
	}
	return sum
}

// END OMIT
