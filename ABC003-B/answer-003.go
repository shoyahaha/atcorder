package main

import (
  "fmt"
)

func main() {
  var S, T string
  fmt.Scan(&S, &T)

  // @で置き換え可能な文字をスライスで定義
  atCoderChars := []rune{'a', 't', 'c', 'o', 'd', 'e', 'r'}

  // 置き換え可能かを判定する関数
  isAtCoderChar := func(c rune) bool {
    for _, char := range atCoderChars {
      if c == char {
        return true
      }
    }
    return false
  }

  canWin := true
  for i := 0; i < len(S); i++ {
    charS := rune(S[i])
    charT := rune(T[i])

    if charS == charT {
      // 両方が同じ文字ならOK
      continue
    }

    if charS == '@' {
      // charSが@の場合、charTが置き換え可能文字かどうか
      if !isAtCoderChar(charT) {
        canWin = false
        break
      }
    } else if charT == '@' {
      // charTが@の場合、charSが置き換え可能文字かどうか
      if !isAtCoderChar(charS) {
        canWin = false
        break
      }
    } else {
      // 両方が異なる場合は負け
      canWin = false
      break
    }
  }

  // 結果を出力
  if canWin {
    fmt.Println("You can win")
  } else {
    fmt.Println("You will lose")
  }
}
