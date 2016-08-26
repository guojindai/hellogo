package main

import "fmt"

func Pic(dx, dy int) [][]uint8 {
  rs := make([][]uint8, dx)
  for i := range rs {
    rs[i] = make([]uint8, dy)
    for j, v := range rs[i] {
      *&v = uint8(i * j)
      fmt.Println(&rs[i][j], v, &v, *&v)
    }
  }
  fmt.Println(rs)
  return rs
}

func main() {
  Pic(4, 4)
}
