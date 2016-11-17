package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9 // 68 / 9 = 7
	y = sum - x     // 17 - 7
	return          // x = 7, y = 10を返すnaked return
}

func main() {
	fmt.Println(split(17))
}
