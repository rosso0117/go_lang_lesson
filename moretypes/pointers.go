package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // iへのポインタ
	fmt.Println(*p) // ポインタの先の変数を指す
	*p = 21         // ポインタの指示す先(i)を21へ
	fmt.Println(i)

	p = &j
	*p = *p / 36
	fmt.Println(j)
}
