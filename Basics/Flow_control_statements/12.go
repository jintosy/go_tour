package main

import "fmt"

func main() {
    // 処理がその関数の最後まで遅延される
    // 評価自体はすぐ行われる
    defer fmt.Println("world")

    fmt.Println("hello")
}
