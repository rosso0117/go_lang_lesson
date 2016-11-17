package main

import "fmt"

func main() {
	sum := 1
	// whie sum < 1000
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
