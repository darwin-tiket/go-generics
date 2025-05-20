package main

// START OMIT
func FindOrdersAsync() (<-chan []Order, <-chan error) {
	outCh := make(chan []Order, 1)
	errCh := make(chan error, 1)

	go func() {
		defer close(outCh)
		defer close(errCh)

		orders, err := findOrders()
		if err != nil {
			errCh <- err
			return
		}
		outCh <- orders
	}()

	return outCh, errCh
}

// END OMIT
