package main

// START OMIT
func Sum(iter interface{}) interface{} {
	switch s := iter.(type) {
	case []int:
		var sum int
		for _, v := range s {
			sum += v
		}
		return sum
	case []int64:
		var sum int64
		for _, v := range s {
			sum += v
		}
		return sum
	case []float64:
		var sum float64
		for _, v := range s {
			sum += v
		}
		return sum
		// []int32, []uint32, []float32, and so on...
	}
}

// END OMIT
