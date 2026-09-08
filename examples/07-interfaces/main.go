package main

import "fmt"

// GreeterはGreetメソッドを持つ型が満たすinterface。
type Greeter interface {
	Greet() string
}

type User struct {
	Name string
}

// UserはGreet() stringを持つため、Greeter interfaceを満たす。
// Goではimplementsのような明示的な宣言は不要。
// 必要なメソッドを実装していれば、自動的にinterfaceを満たす。← Go特有
func (u User) Greet() string {
	return "Hello, " + u.Name
}

// 引数の型としてinterfaceを指定できる。
// gにはGreeterを満たす型の値を渡せる。
// つまり「Greet() stringを持つ値なら型を問わず受け取れる」。
func printGreeting(g Greeter) {
	fmt.Println(g.Greet())
}

func main() {
	user := User{
		Name: "hamru",
	}

	// UserはGreet() stringを持っているため、
	// Greeterを明示的にimplementsしていなくても渡せる。← Go特有
	printGreeting(user)
}
