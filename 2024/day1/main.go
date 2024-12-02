package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func parse(input string) ([]int, []int) {
  input_slice := strings.Split(input, "\n")

  var ls []int
  var rs []int

  for i := range input_slice {
    sp := strings.Fields(input_slice[i])
    l, err := strconv.Atoi(sp[0])
    if err != nil {
      fmt.Println(err.Error())
    }
    r, err := strconv.Atoi(sp[1])
    if err != nil {
      fmt.Println(err.Error())
    }
    ls = append(ls, l)
    rs = append(rs, r)
  }

  return ls, rs
}

func p1(input string) int {
  ls, rs := parse(input)

  sort.Slice(ls, func(i, j int) bool {
    return ls[i] < ls[j]
  })

  sort.Slice(rs, func(i, j int) bool {
    return rs[i] < rs[j]
  })

  sum := 0

  for i := range rs {
    d := rs[i] - ls[i]
    if d < 0 { d = -d }
    sum += d
  }

  return sum
}

func p2(input string) int {
  ls, rs := parse(input)

  sum := 0
  for i := range ls {
    count := 0
    for j := range rs {
      if ls[i] == rs[j] {
        count++
      }
    }
    sum = sum + (count * ls[i])
  }

  return sum
}

func main() {
  fmt.Println(p2(input))
}

