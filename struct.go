package main 

import "fmt"

type Vertex struct {
  x int
  y int
}

var (
  a = 1
  b = 2
)

func main () {
  fmt.Println(Vertex{1, 2})
}