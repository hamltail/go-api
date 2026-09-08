package main

import (
	"errors"
	"fmt"
)

// Goでは例外を投げるのではなく、errorを戻り値として返すことが多い。
func findUser(id int) (string, error) {
	if id <= 0 {
		return "", errors.New("invalid user id")
	}

	if id != 1 {
		return "", errors.New("user not found")
	}

	// エラーがない場合はnilを返す。
	return "hamru", nil
}

func main() {
	name, err := findUser(1)

	// Goではerrorがnilかどうかを呼び出し側で明示的に確認する。
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(name)
}
