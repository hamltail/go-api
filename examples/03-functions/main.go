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
}
