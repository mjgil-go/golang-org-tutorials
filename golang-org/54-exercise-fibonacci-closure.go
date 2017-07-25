package main

import "fmt"

// func recursiveFib(num int) int {
//   if num == 0 {
//     return 0
//   } else if num == 1 {
//     return 1
//   }
//   return recursiveFib(num - 1) + recursiveFib(num - 2)
// }

// fibonacci is a fucntion that returns
// a function that returns an int.
func fibonacci() func() int {
  start, x, y := -1, 0, 1
  returnFn := func() int {
    start++
    if start == 0 {
      return x
    } else if start == 1 {
      return y
    } else {
      xPlusY := x + y
      x = y
      y = xPlusY
      return xPlusY
    }
    panic("unreachable")
  }
  return returnFn
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}
