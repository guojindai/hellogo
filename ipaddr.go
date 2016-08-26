package main

import "fmt"
import "strconv"

type IPAddr [4]byte

func (ia IPAddr) String () string {
  return strconv.Itoa(int(ia[0])) + "." + strconv.Itoa(int(ia[1])) + "." + strconv.Itoa(int(ia[2])) + "." + strconv.Itoa(int(ia[3]))
}

func main() {
  hosts := map[string]IPAddr{
    "loopback":  {127, 0, 0, 1},
    "googleDNS": {8, 8, 8, 8},
  }
  for name, ip := range hosts {
    fmt.Printf("%v: %v\n", name, ip)
  }
}
