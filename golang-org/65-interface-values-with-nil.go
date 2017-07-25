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
func (myString *MyStringStruct) PrintName() {
  if myString == nil {
    fmt.Println("<nil>")
    return
  }
  fmt.Println(myString.Name)
}

func describe(i MyInterface) {
  fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
  var i MyInterface
  var t *MyStringStruct

  i = t
  describe(i)
  i.PrintName()

  i = &MyStringStruct{"Malcom"}
  describe(i)
  i.PrintName()
}
