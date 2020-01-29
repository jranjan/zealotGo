package main

import (
    "fmt"
    "time"
)

func main(){
    c1 := make(chan int)
    c2 := make(chan int)

    go processor_one(c1)
    go processor_two(c2)

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }

    // Let us wait to see all strings printed on our console before
    // we say that we are done.
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}

func processor_one(c chan int){
    fmt.Println("Computing started using processor one...")
    time.Sleep(2 * time.Second)
    c <- 1
    fmt.Println("Done with processor one, notifed as well!")
}

func processor_two(c chan int){
    fmt.Println("Computing started using processor two...")
    time.Sleep(10 * time.Second)
    c <- 2
    fmt.Println("Done with processor two, notifed as well")
}