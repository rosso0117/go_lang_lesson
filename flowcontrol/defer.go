package main

import "fmt"

func main() {
	// 呼び出した関数がreturnするまで実行されない
	defer fmt.Println("world")

	fmt.Println("Hello")
}
