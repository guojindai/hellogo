package main 

import (
  "fmt"
  "math"
)

func Sqrt (x float64) float64 {
  z := 1.0
  pz := 0.0
  for math.Abs(pz - z) > 0.00000001 {
    pz = z
    z = z - (z * z - x) / (z * 2)
    fmt.Println(pz, z)
  }
  return z
}

func main () {
  fmt.Println(Sqrt(2))
  fmt.Println(math.Sqrt(2))
}