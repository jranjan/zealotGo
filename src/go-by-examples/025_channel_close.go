package main

// Closing a channel is basically a hint to receiver that there is nothing
// more to process. It does not mean that channel is flushed out. It does
// not mean that receiver can not fetch the data. By receiver, we mean by
// go routines.

import (
    "math/rand"
    "fmt"
    "time"
)

func main() {
    input := make(chan int, 2)
    done := make(chan bool)

    go requester(input)
    go square_generator(input, done)

    <- done
}

func requester(c chan int){
    for count :=1; count <= 25; count = count +1 {
        random := rand.Int() % 100
        fmt.Printf("Give me square of %d\n", random)
        c <- random
    }
    fmt.Println("Generated too many of job, no more needed. So, closing the channel...")
    close(c)
}

func square_generator(c chan int, done chan bool){
    count := 0
    for {
        number, more := <- c
        if more != true {
            fmt.Printf("No more jobs left...")
            break

        } else {
            square := number * number
            fmt.Printf("Square of %d is %d\n", number, square)

            count = count + 1
            if count % 3 == 0{
                fmt.Printf("taking rest for 6 seconds...\n")
                time.Sleep(time.Second * 6)
                fmt.Println("-----------------------------------------------------------------")
            }
        }
    }

    fmt.Printf("Processed %d jobs\n", count)
    done <- true
}