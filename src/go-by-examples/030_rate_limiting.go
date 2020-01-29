package main

import (
    "time"
    "fmt"
)


// Entire game ia about following activity:
//      1. create a ticker
//      2. ensure that ticker generates a value periodically
//      3. wait for ticker to have value (read the channel)
// Rest will automatically taken care of!
func main() {
    serve_request_every_2s()
    fmt.Println()
    serve_bursty_requests_every_10s()
}


func serve_request_every_2s(){
    fmt.Println("Illustrate the limiter...")
    requests := make(chan int, 10)
    for i := 0; i < 10; i = i + 1 {
        requests <- i
    }
    close(requests)

    time_limiter := time.Tick(time.Second * 2)
    for j := 0; j < 10; j = j + 1 {
        // The below statement pauses the running thread till ticker genereates
        // a new number. Thinks ticker, timer etc as an entity which keeps generating
        // some number. The variables (or data strucuture) can be thought as
        // a channel where ticker/timer entity is posting some input periodically
        // at asked frequency. It is AWESOME WAY of saying that I am not going to
        // serve at frequency higher than I can. Next question comes what will
        // happen to sender. Sender will have to live with followng choices:
        //      1. Use buffer
        //      2. Get blocker till sender request is processed.
        // Awesome, isn't it? A elegant language construct to solve many common
        // programming problems.
        <- time_limiter
        fmt.Printf("Got request id (%d) at %s \n", j, time.Now())
    }
}

func serve_bursty_requests_every_10s(){
    fmt.Println("Illustrate the bursty limiter...")

    burstyRequests := make(chan int, 10)
    for i := 0; i < 10; i = i + 1 {
        burstyRequests <- i
    }
    close(burstyRequests)

    burstyLimiter := make(chan int, 3)

    // This will push three number in channel immediately.
    for i := 0; i < 3; i++ {
        burstyLimiter <- i
    }

    // This will ensure that we keep ticking at every 10 seconds
    // For every tick, we put three values so that three requests gets
    // processed in batch.
    // Did you notice how we did use range? Awesome, isn't it?
    go func() {
        for range time.Tick(time.Second * 10) {
             for i := 0; i < 3; i++ {
                burstyLimiter <- i
             }
        }
    }()


    // Time to serve requests.
    // Serve as fast you can. Since we are limited by by burstyLimiter, we will
    // server three of those quickly but we will get limited to 10 seconds
    // after that. Explaination: three values are ready, and next will come
    // once someone feeds which is another go routine. And feeding happens
    // once in 10s.
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}