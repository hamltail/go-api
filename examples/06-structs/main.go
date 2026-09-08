package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// 値レシーバ
// Userのコピーを受け取る。
func (u User) greet() string {
	return "Hello, " + u.Name
}

// ポインタレシーバ
// 元のUserを変更できる。
func (u *User) updateAge(age int) {
	u.Age = age
}

func main() {
	user := User{
		Name: "hamru",
		Age:  99,
	}

	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Age)

	user.Age = 120
	fmt.Println(user.Age)
	fmt.Println(user.greet())

	// Goが自動的に&を補ってくれるため、(&user).updateAge(...) と書く必要はない。
	user.updateAge(150)
	fmt.Println(user.Age)
}
