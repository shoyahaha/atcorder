package main

import (
  "fmt"
  "strconv"
)

func main() {
  var N, K int
  fmt.Scan(&N, &K)

  for i := 0; i < K; i++ {
	// N が 200 で割り切れる場合は N を 200 で割る
	if N%200 == 0 {
	  N /= 200
	} else {
	  // そうでない場合は、N の末尾に "200" を追加
	  NString := strconv.Itoa(N)
	  NString += "200"
	  N, _ = strconv.Atoi(NString)
	}
  }

  // 結果を出力
  fmt.Println(N)
}
