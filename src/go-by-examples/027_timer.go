package main

import (
    "time"
    "fmt"
)

func main() {
    timer1 := time.NewTimer(time.Second * 2)
    <- timer1.C
    fmt.Printf("Timer 1 expired\n")

    timer2 := time.NewTimer(time.Second * 10)
    go func() {
        <- timer2.C
        fmt.Printf("Timer 2 expired")
    }()
    timer2.Stop()
    fmt.Printf("Timer 2 stopped\n")
}
