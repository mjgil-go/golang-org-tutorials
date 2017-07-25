package main

import (
	"golang.org/x/tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  wordMap := map[string]int{}
  stringArray := strings.Fields(s)
  for i := range stringArray {
    currentString := stringArray[i]
    wordMap[currentString] += 1
  }
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
