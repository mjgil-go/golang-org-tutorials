package main

import "fmt"

func main()  {
  x := 2
  defer fmt.Println(x)

  x = 2*x
  fmt.Println(x)
}
