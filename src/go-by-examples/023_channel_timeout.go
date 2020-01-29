package main

import (
    "fmt"
    "time"
)

func main() {
    done := make(chan bool, 1)

    fmt.Println("There is flaw in code...try to find it...")
    go my_worker(done)

    // There is flaw in calling multiple receivers hooked to same channel. Once one of the receiver
    // fetch the message, queus is empty. So, sender and receiver are blocked. Sender thinks that
    // I am done while receiver keeps waiting for notification/. OK! What is way to implement observer
    // patter using channel? Can we?
    go my_watcherOne(done)
    go my_watcherTwo(done)
    go my_watcherThree(done)


    // Let us wait to see all strings printed on our console before
    // we say that we are done.
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}

func my_worker(signal chan bool){
    fmt.Println("Computing...")
    time.Sleep(10 * time.Second)
    fmt.Println("Computation done, notifying receivers...")
    signal <- true
}

func my_watcherOne(signal chan bool){
    fmt.Println("Started watcherOne")
    select {
        case <- signal:
            fmt.Println("my_watcherOne receieved data")
        case <-time.After(time.Second * 12):
            fmt.Println("my_watcherOne timeout after 15 seconds...")
    }
}

func my_watcherTwo(signal chan bool){
    fmt.Println("Started watcherTwo")
    select {
        case res := <- signal:
            fmt.Println("my_watcherTwo receieved data")
            fmt.Println(res)
        case <-time.After(time.Second * 15):
            fmt.Println("my_watcherTwo timeout after 15 seconds...")
    }
}

func my_watcherThree(signal chan bool){
    fmt.Println("Started watcherThree")
    select {
        case res := <- signal:
            fmt.Println("my_watcherThree receieved data")
            fmt.Println(res)
        case <-time.After(time.Second * 20):
            fmt.Println("my_watcherThree timeout after 20 seconds...")
    }
}


