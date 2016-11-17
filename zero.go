package main

import "fmt"

// 初期値を与えないと、ゼロ値が与えられる
func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
