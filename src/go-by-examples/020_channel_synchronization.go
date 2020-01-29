package main

import (
    "fmt"
    "time"
)

func main() {
    done := make(chan bool, 1)

    fmt.Println("There is flaw in code...try to find it...")
    go worker(done)

    // There is flaw in calling multiple receivers hooked to same channel. Once one of the receiver
    // fetch the message, queus is empty. So, sender and receiver are blocked. Sender thinks that
    // I am done while receiver keeps waiting for notification/. OK! What is way to implement observer
    // patter using channel? Can we?
    go watcherOne(done)
    go watcherTwo(done)
    go watcherThree(done)


    // Let us wait to see all strings printed on our console before
    // we say that we are done.
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}

func worker(signal chan bool){
    fmt.Println("Computing...")
    time.Sleep(30 * time.Second)
    fmt.Println("Computation done, notifying receivers...")
    signal <- true
}

func watcherOne(signal chan bool){
    fmt.Println("Started watcherOne")
    <- signal
    fmt.Println("watcherOne receieved data")
}

func watcherTwo(signal chan bool){
    fmt.Println("Started watcherTwo")
    <- signal
    fmt.Println("watcherTwo receieved data")
}

func watcherThree(signal chan bool){
    fmt.Println("Started watcherThree")
    <- signal
    fmt.Println("watcherThree receieved data")
}

