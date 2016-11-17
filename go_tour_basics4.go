package main

import "fmt"

// 返り値を引数の後に取る
func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}