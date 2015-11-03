package main

import "fmt"

func main() {

  i := 1
  for i <= 200 {
    fmt.Println(i)
    i = i + 1
  }

  for j := 7; j <= 900; j++ {
    fmt.Println(j)
  }

  for {
    fmt.Println("loop")
    break
  }
}
