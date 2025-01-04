package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scan(&n)

  switch {
  case n >= 0 && n < 100:
	fmt.Println("00")
  case n >= 100 && n <= 5000:
	fmt.Printf("%02d\n", n/100)
  case n >= 6000 && n <= 30000:
	fmt.Println(n/1000 + 50)
  case n >= 35000 && n <= 70000:
    fmt.Println((n/1000-30)/5 + 80)
  default:
	fmt.Println("89")
  }
}
