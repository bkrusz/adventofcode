package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parse(input string) ([][]int) {
  input_slice := strings.Split(input, "\n")

  var sli [][]int
  for i := range input_slice {
    split := strings.Split(input_slice[i], " ")
    var is []int
    for j := range split {
      in, err := strconv.Atoi(split[j])
      if err != nil {
        fmt.Println("Error converting to int")
      }
      is = append(is, in)
    }
    sli = append(sli, is)
  }

  return sli
}

func p1(input string) int {
  list := parse(input)

  sum := 0
  for i := 0; i < len(list); i++ {
    if list[i][0] > list[i][1] && list[i][0] - list[i][1] < 4 {
      v := true
      for j := 0; j < len(list[i])-1; j++ {
        if !(list[i][j] > list[i][j+1]) || !(list[i][j] - list[i][j+1] < 4){
          v = false
        }
      }
      if v { sum++ }
    }
    if list[i][1] > list[i][0] && list[i][1] - list[i][0] < 4 {
      v := true
      for j := 0; j < len(list[i])-1; j++ {
        if !(list[i][j+1] > list[i][j]) || !(list[i][j+1] - list[i][j] < 4){
          v = false
        }
      }
      if v { sum++ }
    }
  }

  return sum
}

func p2(input string) int {
  sum := 0

  return sum
}

func main() {
  fmt.Println(p1(input))
}

