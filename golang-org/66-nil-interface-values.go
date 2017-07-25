package main

import "fmt"

type MyInterface interface {
  PrintName()
}

func describe(i MyInterface) {
  fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
  var i MyInterface
  describe(i)
  i.PrintName() // runtime error because no underlying type
}
