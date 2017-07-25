package main

// run "go get golang.org/x/tour/pic" in the terminal
// need to run on the website to see image
// https://tour.golang.org/moretypes/18
import (
  "fmt"
)

func Pic(dx, dy int) [][]uint8 {
  returnVal := [][]uint8{}
  for j := 0; j < dy; j++ {
    xArray := []uint8{}
    for i := 0; i < dx; i++ {
        xArray = append(xArray, uint8(i*j))
        // xArray = append(xArray, uint8(i^j))
        // xArray = append(xArray, uint8((i+j)/2))
    }
    returnVal = append(returnVal, xArray)
  }
  return returnVal
}


func Pic2(dx, dy int) [][]uint8 {
  returnVal := make([][]uint8, dy)
  for j := range returnVal {
    xArray := make([]uint8, dx)
    for i := range xArray {
      xArray[i] = uint8(i*j)
      // xArray[i] = uint8(i^j)
      // xArray[i] = uint8((i+j)/2))
    }
    returnVal[j] = xArray
  }
  return returnVal
}

func main() {
  // pic.Show(Pic2)
  fmt.Println(Pic(5, 3))
  fmt.Println(Pic2(5, 3))
}
