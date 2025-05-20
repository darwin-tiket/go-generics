package main

// START OMIT
func Foo[T interface{ ~int }]() {}

func Bar[U interface{ int | string }]() {}

// END OMIT
