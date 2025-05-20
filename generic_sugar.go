package main

// START OMIT
func Foo[T ~int]() {}

func Bar[U int | string]() {}

// END OMIT
