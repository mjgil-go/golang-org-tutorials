package main

import "fmt"

func main()  {
  s := []int{1, 2, 3, 4, 5, 6}
  printSlice(s)

  // Slice the slice to give it zero length
  s = s[:0]
  printSlice(s)

  // Extend its length
  s = s[:4]
  printSlice(s)

  // Drop its first two values
  s = s[2:]
  printSlice(s)
}

func printSlice(s []int)  {
  fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
