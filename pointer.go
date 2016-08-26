package main

import "fmt"

func main () {
  i := 1
  p := &i
  fmt.Println(i, *p)
  *p = 2
  fmt.Println(i, *p)
}