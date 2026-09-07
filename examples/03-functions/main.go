package main

import (
	"errors"
	"fmt"
)

func greet(name string) string {
	return "Hello, " + name
}

func add(a int, b int) int {
	return a + b
}

// Goの関数は複数の値を返せる。
func divide(a int, b int) (int, int) {
	return a / b, a % b
}

// Goでは処理結果とerrorを複数の戻り値として返すことが多い。
func validateAge(age int) (string, error) {
	if age < 0 {
		return "", errors.New("age must be 0 or greater")
	}

	return "valid age", nil
}

// 戻り値には名前を付けることもできる。
func calculate(a int, b int) (sum int, difference int) {
	sum = a + b
	difference = a - b

	return sum, difference
}

// ... を使うと、同じ型の引数を可変個受け取れる。
func sumAll(numbers ...int) int {
	total := 0

	// Goではforがループ構文を担い、foreachやwhileの専用キーワードはない。
	// rangeは「インデックス, 値」を返すが、インデックスは不要なので_で明示的に捨てる。
	// これはシンプルでいいね。
	for _, number := range numbers {
		total += number
	}

	return total
}

// 関数を引数として受け取ることもできる。
// operationには、呼び出し側から渡されたコールバック関数が入る。
func executeCalculation(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func main() {
	message := greet("hamru")
	fmt.Println(message)

	result := add(10, 20)
	fmt.Println(result)

	quotient, remainder := divide(10, 3)
	fmt.Println(quotient)
	fmt.Println(remainder)

	message, err := validateAge(99)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(message)
	}

	sum, difference := calculate(10, 3)
	fmt.Println(sum)
	fmt.Println(difference)

	total := sumAll(10, 20, 30, 40)
	fmt.Println(total)

	// 関数の代入
	addFunction := add
	functionResult := addFunction(10, 20)
	fmt.Println(functionResult)

	// Goにはアロー関数の構文はないが、無名関数を定義できる。
	multiply := func(a int, b int) int {
		return a * b
	}
	multiplyResult := multiply(10, 20)
	fmt.Println(multiplyResult)

	// add関数をコールバックとして渡す。
	calculationResult := executeCalculation(10, 20, add)
	fmt.Println(calculationResult)
}
