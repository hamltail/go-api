// 実行:
//   go run ./examples/01-basics
//
// ビルド:
//   go build -o bin/01-basics ./examples/01-basics
//
// ビルド後の実行:
//   ./bin/01-basics

// package main は実行可能なプログラムを定義する特別なパッケージ
package main

// fmt（format）は文字列の整形や標準入出力を扱う標準パッケージ
import "fmt"

// main 関数はプログラムのエントリーポイント
func main() {
	fmt.Println("Hello, Go!")
}
