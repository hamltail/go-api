package main

import "fmt"

func main() {
	var name string = "hamru"
	fmt.Println(name)

	// 型推論
	var age = 20
	fmt.Println(age)

	// 短縮変数宣言
	// := は変数の宣言と初期化を短く書く構文。
	// 関数の中でのみ使用できる。
	language := "Go"
	fmt.Println(language)

	// 定数
	const version = "1.0.0"
	fmt.Println(version)

	// 基本型
	var message string = "Hello"
	var count int = 10
	var rate float64 = 1.5
	var enabled bool = true

	fmt.Println(message)
	fmt.Println(count)
	fmt.Println(rate)
	fmt.Println(enabled)

	// 初期値を指定しない変数には、型ごとのゼロ値が設定される。
	var zeroString string
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool

	fmt.Println(zeroString)
	fmt.Println(zeroInt)
	fmt.Println(zeroFloat)
	fmt.Println(zeroBool)

	// := では複数の変数を同時に宣言できる。
	x, y := 10, 20
	fmt.Println(x, y)

	// エラー
	// x, y := 30, 40

	// 少なくとも1つ新しい変数があれば、宣言済みの変数も一緒に代入できる。
	// これは嫌だなぁ
	x, z := 30, 40
	fmt.Println(x, z)

	// value, err := something()
	// 　からの
	// otherValue, err := somethingElse() といった形で遭遇するらしい。

	// 異なる型の値は明示的に型変換する。Goは暗黙的に変換しない。安全なキャストでもNG
	var number int = 10
	var decimal float64 = float64(number)

	fmt.Println(number)
	fmt.Println(decimal)

	// float64をintに変換すると、小数部分は切り捨てられる。
	var price float64 = 19.8
	var integerPrice int = int(price)

	fmt.Println(price)
	fmt.Println(integerPrice)
}
