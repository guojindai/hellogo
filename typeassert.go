package main

import "fmt"

func main() {
  var i interface{} = "hello"

  s, ok := i.(string)
  fmt.Println(s, ok, i)
  s = "211"
  fmt.Println(s, ok, i)

  s2 := &i
  *s2 = "321"
  fmt.Println(*s2, ok, i)
}
