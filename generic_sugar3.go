package main

// START OMIT
type MyType interface {
	int | string
}

func Foo[T MyType]() {}

// END OMIT
