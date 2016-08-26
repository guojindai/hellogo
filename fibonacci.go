package main

import "fmt"

func fibonacci() func() int {
  i, v1, v2 := 0, 0, 1
  return func () int {
    var v int
    if i == 0 || i == 1 {
      v = i
    } else {
      v = v1 + v2
      v1 = v2
      v2 = v
    }
    i ++
    return v
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 10; i++ {
    fmt.Println(f())
  }
}
