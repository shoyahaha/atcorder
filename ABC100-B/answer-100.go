package main

import "fmt"

func main() {
  var D, N int
  fmt.Scan(&D, &N)

  count := 0
  for i := 1; ; i++ {
    num := i
    divisions := 0

    // 100で割り切れる回数を計算
    for num%100 == 0 {
      divisions++
      num /= 100
    }

    // ちょうどD回割り切れる場合
    if divisions == D {
      count++
    }

    // N番目に到達したら出力
    if count == N {
      fmt.Println(i)
      break
    }
  }
}
