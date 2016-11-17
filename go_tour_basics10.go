package main

import "fmt"

func main() {
    var i, j int = 1, 2
    // 関数内ではvarを使わなくても暗黙型を使える
    k := 3
    c, python, java := true, false, "no!"

    fmt.Println(i, j, k, c, python, java)
}