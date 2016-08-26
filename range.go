package main 

import "fmt"

var pow = []int{11, 22, 33}


func main () {
  for i, v := range pow {
    fmt.Println(i, v)
  }
}