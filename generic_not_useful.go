package main

import "io"

// START OMIT
// Good
func Foo(r io.Reader) {
	r.Read()
}

// Bad
func Foo[R io.Reader](r R) {
	r.Read()
}

// END OMIT
