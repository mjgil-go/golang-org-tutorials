package main

// run "go get golang.org/x/tour/pic" in the terminal
// need to run on the website to see image
// https://tour.golang.org/moretypes/18
import (
  "golang.org/x/tour/pic"
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

func main() {
  pic.Show(Pic)
  // fmt.Println(Pic(5, 3))
}
