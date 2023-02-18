package main

import "fmt"

func main() {
    fmt.Println("counting")

    // deferで遅延実行する処理はスタックされる
    // なのでFILOになる
    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }

    fmt.Println("done")
}
