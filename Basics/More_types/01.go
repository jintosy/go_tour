package main

import "fmt"

func main() {
    i, j := 42, 2701

    p := &i  // iへのポインタを作成
    fmt.Println(*p) // ポインタを参照

    *p = 21  // ポインタを通じてiの値を変更
    fmt.Println(i)

    p = &j
    *p /= 37
    fmt.Println(j)
}
