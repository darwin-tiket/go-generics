package main

// START OMIT
   Type parameter
         ^
         |
func Foo[T any]() {}
            |
            v
      Type constraint


v := Foo[int]()
          |
          v
   Type argument
// END OMIT
