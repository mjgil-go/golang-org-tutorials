package main

import (
  "fmt"
  "math"
)

type MyInterface interface {
  PrintName()
}

type MyStringStruct struct {
  Name string
}

// this means MyStringStruct implements MyInterface
// but we don't need to explicitly declare it
func (myString *MyStringStruct) PrintName() {
  fmt.Println(myString.Name)
}

type MyFloat float64

func (myFloat MyFloat) PrintName() {
  fmt.Println(myFloat)
}

func describe(i MyInterface) {
  fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
  var i MyInterface

  i = &MyStringStruct{"Malcom"}
  describe(i)
  i.PrintName()

  i = MyFloat(math.Pi)
  describe(i)
  i.PrintName()
}
