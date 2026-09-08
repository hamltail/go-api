package main

import "fmt"

func main() {
	// arrayは要素数が型の一部になる。
	numbersArray := [3]int{10, 20, 30}
	fmt.Println(numbersArray)

	// sliceは要素数を型に含まない。
	// Goでは固定長のarrayより、sliceを使うことが多い。
	numbersSlice := []int{10, 20, 30}
	fmt.Println(numbersSlice)

	// appendでsliceの末尾に要素を追加する。
	// appendは元のsliceを直接変更するのではなく、追加後のsliceを返す。
	numbersSlice = append(numbersSlice, 40)
	fmt.Println(numbersSlice)
	numbersSlice = append(numbersSlice, 50, 60)
	fmt.Println(numbersSlice)

	fmt.Println("len:", len(numbersSlice))
	// capはsliceが現在確保している領域で保持できる要素数（容量）を返す。
	// appendで容量が足りなくなると、Goがより大きな領域を確保してsliceを成長させる。
	// 通常はGoが管理してくれるため、基本的な利用ではcapを意識する必要はない。
	fmt.Println("cap:", cap(numbersSlice))

	// インデックスを指定して要素を取得する。
	fmt.Println(numbersSlice[0])
	fmt.Println(numbersSlice[2])

	// [開始位置:終了位置] でsliceの一部分を切り出せる。
	// 開始位置は含み、終了位置は含まない。
	part := numbersSlice[1:4]
	fmt.Println(part)

	// 切り出したsliceは、元のsliceと同じデータを参照している。
	// そのため、partの要素を変更するとnumbersSlice側にも反映される。
	part[0] = 999
	fmt.Println(part)
	fmt.Println(numbersSlice)

	// makeを使うと、長さを指定してsliceを作成できる。
	// 各要素は型のゼロ値で初期化される。
	madeSlice := make([]int, 3)

	fmt.Println(madeSlice)
	fmt.Println("len:", len(madeSlice))
	fmt.Println("cap:", cap(madeSlice))

	// 第3引数を指定すると、長さとは別に容量を指定できる。
	reservedSlice := make([]int, 3, 10)

	fmt.Println(reservedSlice)
	fmt.Println("len:", len(reservedSlice))
	fmt.Println("cap:", cap(reservedSlice))

	// mapはキーと値の組を保持する。
	// map[キーの型]値の型 という形で型を表す。
	scores := map[string]int{
		"hamru":  100,
		"felina": 99,
	}

	fmt.Println(scores)

	// キーを指定して値を取得する。
	fmt.Println(scores["hamru"])
	// 新しいキーと値を追加する。
	scores["seek"] = 80
	// 既存のキーを指定すると値を更新する。
	scores["hamru"] = 120

	fmt.Println(scores)

	// 存在しないキーを取得すると、値の型のゼロ値が返る。
	// そのため、値だけでは「キーが存在しない」のか「ゼロ値が登録されている」のか区別できない。
	fmt.Println(scores["unknown"])

	// 2つ目の戻り値でキーが存在するか確認できる。
	// 慣例的にokという変数名を使うことが多い。
	score, ok := scores["hamru"]

	fmt.Println(score)
	fmt.Println(ok)

	// deleteで指定したキーと値をmapから削除する。
	delete(scores, "seek")
	fmt.Println(scores)

	// rangeを使ってmapのキーと値を順番に取得できる。
	// mapの走査順序は保証されない。
	for name, score := range scores {
		fmt.Println(name, score)
	}

	// makeを使って、要素を後から追加できる空のmapを作成する。
	users := make(map[string]int)

	users["alice"] = 20
	users["bob"] = 30

	fmt.Println(users)
}
