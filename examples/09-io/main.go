package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// struct tagを使って、GoのフィールドとJSONのキーを対応付ける。
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// ファイル全体を[]byteとして読み込む。
	// os.ReadFileも「値 + error」を返すGoの定番パターン。
	data, err := os.ReadFile("examples/09-io/data.json")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// []byteをstringに変換して表示する。
	fmt.Println(string(data))

	var user User

	// JSONの[]byteをGoのstructに変換する。
	// Unmarshalがuserを書き換えられるように、&でuserのアドレスを渡す。
	// 通常の関数呼び出しなので、メソッドのポインタレシーバと違って&は自動補完されない。
	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Age)
}
