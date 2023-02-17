package main

import (
    "fmt"
)

func main() {
    // Type conversions
    /*
    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(x*x + y*y))
    var z uint = uint(f)

    fmt.Println(x, y, z)*/

    // more simple
    i := 42
    f := float64(i)
    u := uint(f)
    fmt.Println(i, f, u)
}
