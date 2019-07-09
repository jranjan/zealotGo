package main

import (
    "fmt"
)

func main() {
    for i := 1; i < 10; i = i + 1 {
        myno := i
        squareme(&myno)
        fmt.Printf("Square of %d is %d\n", i, myno)
    }
}

func squareme(num_p *int){
    x := *num_p
    square := x * x
    *num_p = square
}