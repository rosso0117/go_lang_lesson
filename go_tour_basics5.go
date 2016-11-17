package main

import "fmt"

// 引数が同じ型の場合は省略可能
func add(x, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}