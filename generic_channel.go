package main

// START OMIT
func DoAsync[T any](f func() (T, error)) (<-chan T, <-chan error) {
	outCh := make(chan T, 1)
	errCh := make(chan error, 1)

	go func() {
		defer close(outCh)
		defer close(errCh)

		v, err := f()
		if err != nil {
			errCh <- err
			return
		}
		outCh <- v
	}()

	return outCh, errCh
}

// END OMIT
