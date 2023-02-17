package main

import "fmt"

const Pi = 3.14

func main() {
    // Constants
    // only use char,string,bool,numeric
    const World = "世界"
    
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")

    const Truth = true
    fmt.Println("Go rules?", Truth)
}
