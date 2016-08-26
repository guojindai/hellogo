package main

import (
  "io"
  "os"
  "strings"
)

type rot13Reader struct {
  r io.Reader
}

func (rotReader rot13Reader) Read (bs []byte) (int, error) {
  n, err := rotReader.r.Read(bs)
  if err == io.EOF {
    return 0, io.EOF
  } else {
    for i := 0; i < n; i ++ {
      bs[i] += 13
    }
    return n, err
  }
}

func main() {
  s := strings.NewReader("abc")
  r := rot13Reader{s}
  io.Copy(os.Stdout, &r)
}
