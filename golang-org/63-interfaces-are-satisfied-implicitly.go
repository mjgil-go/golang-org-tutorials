package main

import "fmt"

type MyInterface interface {
  PrintName()
}

type MyStringStruct struct {
  Name string
}

// this means MyStringStruct implements MyInterface
// but we don't need to explicitly declare it
func (myString MyStringStruct) PrintName() {
  fmt.Println(myString.Name)
}

func main() {
  var i MyInterface = MyStringStruct{"Malcom"}
  i.PrintName()
}
