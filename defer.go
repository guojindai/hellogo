package main

import "fmt"

func main() {
  
  defer f ()
  fmt.Println("3")
}

func f () {
  defer fmt.Println("1")
  defer fmt.Println("4")
  fmt.Println("2")
}