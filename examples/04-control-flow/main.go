package main

import "fmt"

func main() {
	age := 20

	if age >= 65 {
		fmt.Println("senior")
	} else if age >= 18 {
		fmt.Println("adult")
	} else {
		fmt.Println("minor")
	}

	// ifの初期化文
	if score := 80; score >= 60 {
		fmt.Println("pass")
	} else {
		fmt.Println("fail")
	}

	// switch
	day := "Sunday"

	// Goでは各caseの最後にbreakを書く必要がない。
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("weekend")
	case "Monday":
		fmt.Println("start of week")
	default:
		fmt.Println("weekday")
	}

	// Goではループ処理をforで表現する。
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Goにはwhileキーワードがなく、条件だけのforで同じ処理を表現する。
	count := 0

	for count < 3 {
		fmt.Println(count)
		count++
	}

	// 条件を書かないforは無限ループになる。
	number := 0

	for {
		number++

		if number == 3 {
			break
		}

		fmt.Println(number)
	}

	// rangeを使うと、コレクションの要素を順番に処理できる。
	// rangeは関数ではなく、forと組み合わせて使用するGoの構文。
	numbers := []int{10, 20, 30}

	for index, value := range numbers {
		fmt.Println(index, value)
	}

	// continue
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}

		fmt.Println(i)
	}
}
