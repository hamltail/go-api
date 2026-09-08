package main

import (
	"fmt"
	"time"
)

func makeCoffee(coffeeDone chan bool) {
	for i := 1; i <= 3; i++ {
		fmt.Println("☕ コーヒー", i, "杯目を作っています")
		time.Sleep(300 * time.Millisecond)
	}

	fmt.Println("☕ コーヒー3杯完成！")
	coffeeDone <- true
}

func makeToast(toastDone chan bool) {
	for i := 1; i <= 2; i++ {
		fmt.Println("🍞 トースト", i, "枚目を焼いています")
		time.Sleep(1 * time.Second)
	}

	fmt.Println("🍞 トースト2枚完成！")
	toastDone <- true
}

func main() {
	fmt.Println("🐾 Animal Cafeへようこそ！")
	fmt.Println("ご注文：コーヒー3杯、トースト2枚")
	fmt.Println("------------------------------")

	coffeeDone := make(chan bool)
	toastDone := make(chan bool)

	go makeCoffee(coffeeDone)
	go makeToast(toastDone)

	<-coffeeDone
	<-toastDone

	fmt.Println("------------------------------")
	fmt.Println("☕🍞 ご注文がすべて完成しました！")
	fmt.Println("ごゆっくりどうぞ 🐱")
}
