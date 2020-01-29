package main

import (
    "time"
    "fmt"
)

func main(){
    test_channel()
}

func test_channel(){
    data := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }
    n := len(data)
    c1 := make(chan int)
    c2 := make(chan int)

    go adder(data[:n/2], c1)
    go adder(data[n/2:], c2)

    select {
    case r1 := <- c1:
        fmt.Println(r1)
    case r2 := <- c2:
        fmt.Println(r2)
    case t := <- time.After(3*time.Second):
        fmt.Println(t)
    }
}


func adder(a []int, c chan int){
    s := 0
    for _, x := range a {
        //time.Sleep(time.Second)
        s = s + x
    }
    c <- s
}

