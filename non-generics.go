package main

// START OMIT
func SumInt(iter []int) int {
	var sum int
	for _, v := range iter {
		sum += v
	}
	return sum
}

func SumFloat64(iter []float64) float64 {
	var sum float64
	for _, v := range iter {
		sum += v
	}
	return sum
}

// SumInt32, SumInt64, SumFloat32, and so on...
// END OMIT
