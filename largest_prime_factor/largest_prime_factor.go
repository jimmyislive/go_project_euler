package main

import (
    "fmt"
    "math"
)

func isFactorPrime(factor float64) bool {
    //first check if even
    if math.Mod(factor, 2.0) == 0 {
        return false
    }

    for i:=2; i<int(math.Sqrt(factor)); i++ {
        if math.Mod(factor, float64(i)) == 0 {
            return false
        }
    }
    return true
}

func main() {
    largest := 0
    c := make(chan int)
    go func(channel chan<- int) {
        for i:=2; i<int(math.Sqrt(float64(600851475143))); i++ {

            if math.Mod(600851475143, float64(i)) == 0 {
                if isFactorPrime(float64(i)) == true {
                    channel <- i
                }
            }
        }
        close(channel)
    }(c)

    for num:= range c {
        if largest < num {
            largest = num
        }
    }

    fmt.Println(largest)
}
