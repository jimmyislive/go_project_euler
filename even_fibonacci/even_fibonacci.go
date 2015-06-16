package main

import (
    "fmt"
    "math"
)

func generate(c chan<- int) {
    prev := 1
    current := 2
    c <- prev

    for current <= 4000000 {
        c <- current
        current, prev = current + prev, current
    }
    close(c)
}

func main() {
    sum := 0
    c := make(chan int)
    go generate(c)
    for num := range c {
        if math.Mod(float64(num), 2) == 0 {
            sum = sum + num
        }
    }

    fmt.Println(sum)
}
