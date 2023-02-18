package main


func main() {
    sum := 1
    // equal while
    for sum < 1000 {
        sum += sum
    }

    println(sum)
}
