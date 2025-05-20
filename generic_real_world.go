package main

// START OMIT
func SliceMap[T, V any](iter []T, f func(v T) V) []V {
	out := make([]V, len(iter))
	for _, v := range iter {
		out[i] = f(v)
	}
	return out
}

// END OMIT
