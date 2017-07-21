package main

import (
  "fmt"
  "math"
)

func Sqrt(x float64) float64 {
  z := x
  for y := 10; y > 0; y-- {
    top := math.Pow(z, 2) - x
    bottom := 2 * z
    // fmt.Println("top is: ", top)
    // fmt.Println("bottom is: ", bottom)
    z = z - top/bottom
    // fmt.Println("z is: ", z)
    // fmt.Println("x is: ", x)
  }
  return z
}

func main()  {
  fmt.Println(Sqrt(2))
  fmt.Println(math.Sqrt(2))
}
