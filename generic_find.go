package main

// START OMIT
func SliceAny[T comparable](items []T, expectedItem T) bool {
	for _, item := range items {
		if item == expectedItem {
			return true
		}
	}
	return false
}

// END OMIT
