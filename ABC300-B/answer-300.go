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
  sc = bufio.NewScanner(stdin)

  sc.Split(bufio.ScanWords)

  h, w := readInt(), readInt()

  // グリッドAを保持するスライス
  as := make([]string, h)
  for i := range as {
    as[i] = readString()
  }

  // グリッドBを保持するスライス
  bs := make([]string, h)
  for i := range bs {
    bs[i] = readString()
  }
  // 列方向にi回シフトするループ
  for i := 0; i < w; i++ {
	// 行方向にj回シフトするループ
    for j := 0; j < h; j++ {
      flg := true
	  // グリッドBの各行をループ
      for bi, b := range bs {
        idx := (bi + j) % h
        a := as[idx][i:] + as[idx][:i]
		// グリッドAとグリッドBが一致しない場合
        if b != a {
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
