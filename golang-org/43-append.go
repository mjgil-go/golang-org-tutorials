package main

import "fmt"

func main()  {
  var s []int
  printSlice(s)

  // append works on nil slices
  s = append(s, 0)
  printSlice(s)

  // the slice grows as needed.
  s = append(s, 1)
  printSlice(s)

  // We can add more than one element at a time
  s = append(s, 2, 3, 4) // cap=6 for some reason when appending multiple
  printSlice(s)

  var v []int
  v = append(v, 0, 1, 2, 3, 4)
  printSlice(v)

  x := []int{0, 1, 2, 3, 4}
  printSlice(x)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
