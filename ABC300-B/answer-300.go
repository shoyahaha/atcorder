package main

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "strconv"
)

// bufio.Scannerをグローバル変数として定義
var sc *bufio.Scanner

// 標準入力から整数を読み取る関数
func readInt() int {
  sc.Scan()
  i, _ := strconv.Atoi(sc.Text())
  return i
}

// 標準入力から文字列を読み取る関数
func readString() string {
  sc.Scan()
  return sc.Text()
}

// メイン処理を実行する関数
func run(stdin io.Reader, out io.Writer) {
  sc = bufio.NewScanner(stdin) // 標準入力をスキャナーにセット
  sc.Split(bufio.ScanWords)    // スキャナーを単語単位で分割

  h, w := readInt(), readInt()

  as := make([]string, h) // グリッドAを保持するスライス
  for i := range as {
    as[i] = readString()
  }

  bs := make([]string, h) // グリッドBを保持するスライス
  for i := range bs {
    bs[i] = readString()
  }

  for i := 0; i < w; i++ { // 列方向にi回シフトするループ
    for j := 0; j < h; j++ { // 行方向にj回シフトするループ
      flg := true
      for bi, b := range bs { // グリッドBの各行をループ
        idx := (bi + j) % h
        a := as[idx][i:] + as[idx][:i]
        if b != a { // グリッドAとグリッドBが一致しない場合
          flg = false
          break
        }
      }
      if flg {
        fmt.Fprint(out, "Yes")
        return
      }
    }
  }

  fmt.Fprint(out, "No")
}

// 標準入力と標準出力を渡して実行
func main() {
  run(os.Stdin, os.Stdout)
}
