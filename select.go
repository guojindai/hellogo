package main

import "fmt"

func fibonacci(c chan int) {
  x, y := 0, 1
  for {
    select {
    case c <- x:
      if x > 100 {
        fmt.Println("quit")
        return
      } else {
        x, y = y, x+y
      }
    }
  }
}

func main() {
  c := make(chan int)
  go func() {
    for i := range c {
      fmt.Println(i)
    }
  }()
  fibonacci(c)
}
