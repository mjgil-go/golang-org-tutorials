package main

import "fmt"

func recursiveFib(num int) int {
  if num == 0 {
    return 0
  } else if num == 1 {
    return 1
  }
  return recursiveFib(num - 1) + recursiveFib(num - 2)
}
// fibonacci is a fucntion that returns
// a function that returns an int.
func fibonacci() func() int {
  start := -1
  returnFn := func() int {
    start++
    return recursiveFib(start)
  }
  return returnFn
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}
