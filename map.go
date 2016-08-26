package main

import (
  "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
  m := make(map[string]int)
  for i, w := range strings.Fields(s) {
    m[w] ++
  }
  return m
}

func main() {
  wc.Test(WordCount)
}
