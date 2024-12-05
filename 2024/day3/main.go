package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func p1(input string) int {
  regex := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)")
  mul := regex.FindAllString(input, -1)

  sum := 0
  for m := range mul {
    is := strings.Split(strings.TrimSuffix(strings.TrimPrefix(mul[m],"mul("),")"), ",")
    m1, err := strconv.Atoi(is[0])
    if err != nil { fmt.Printf("Error formatting string \"%v\"\n", is[0]) }
    m2, err := strconv.Atoi(is[1])
    if err != nil { fmt.Printf("Error formatting string \"%v\"\n", is[1]) }
    sum = sum + (m1 * m2)
  }

  return sum
}

func p2(input string) int {
  regex := regexp.MustCompile("mul\\(\\d{1,3},\\d{1,3}\\)|don't\\(\\)|do\\(\\)")
  mul := regex.FindAllString(input, -1)

  sum := 0
  enabled := true
  for m := range mul {
    if mul[m] == "don't()" {
      enabled = false
    }
    if mul[m] == "do()" {
      enabled = true 
    }
    if enabled && strings.Contains(mul[m], "mul") {
      is := strings.Split(strings.TrimSuffix(strings.TrimPrefix(mul[m],"mul("),")"), ",")
      m1, err := strconv.Atoi(is[0])
      if err != nil { fmt.Printf("Error formatting string \"%v\"\n", is[0]) }
      m2, err := strconv.Atoi(is[1])
      if err != nil { fmt.Printf("Error formatting string \"%v\"\n", is[1]) }
      sum = sum + (m1 * m2)
    }
  }

  return sum
}

func main() {
  fmt.Println(p2(input))
}

