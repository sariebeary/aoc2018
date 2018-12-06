// Advent of Code 2018
// Day 2: Inventory Management System
package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

func getTwosThrees(boxID string)(two, three int) {
  count := make(map[rune]int)
  for _, letter := range boxID {
    count[letter] = count[letter] + 1
  }
  for _, freq := range count {
    if freq == 2 {
      two = 1
    } else if freq == 3 {
      three = 1
    }
  }
  return
}

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner:= bufio.NewScanner(file)
  var two, three int
  for scanner.Scan() {
    boxID := scanner.Text()
    x, y := getTwosThrees(boxID)
    two = two + x
    three = three + y
  }
  fmt.Println(two * three)

}
