package main

import (
  "fmt"
  "math/rand"
)




func main() {
  var numbers [100]int

  for i :=0; i <100 ; i++{
    numbers[i] = rand.Intn(1000)
  }

  fmt.Println(numbers)

}