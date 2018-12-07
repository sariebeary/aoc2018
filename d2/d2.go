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

func getCommonID(boxIDs []string) (string){
  for i := 0; i < len(boxIDs); i++ {
    for j := i+1; j < len(boxIDs); j++ {
      differences := 0
      matches := make([]rune, 0)
      box1 := []rune(boxIDs[i])
      box2 := []rune(boxIDs[j])
      for k,_ := range box1 {
        if box1[k] != box2[k] {
          if differences > 1 {
            break
          } else {
            differences++
          }
        } else {
          matches = append(matches, box1[k])
        }
      }
      if differences == 1 {
        return string(matches)
      }
    }
  }
  return ""
}

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  scanner:= bufio.NewScanner(file)
  var two, three int
  var boxIDs []string
  for scanner.Scan() {
    boxID := scanner.Text()
    boxIDs = append(boxIDs, boxID)
    x, y := getTwosThrees(boxID)
    two = two + x
    three = three + y
  }

  fmt.Println("Part 1 Checksum: ", two * three)

  fmt.Println("Part 2 Common ID: ", getCommonID(boxIDs))

}
