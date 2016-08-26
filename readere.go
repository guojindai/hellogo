package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(bs []byte) (int, error) {
  ab := []byte("A")[0]
  for i := range bs {
    bs[i] = ab
  }
  return len(bs), nil
}


func main() {
  reader.Validate(MyReader{})
}
