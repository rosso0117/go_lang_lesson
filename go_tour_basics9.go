package main

import "fmt"

// 初期化子
var i, j int = 1, 2

func main() {
	// 型省略が可能
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
