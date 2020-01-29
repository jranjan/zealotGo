package main

import (
    "fmt"
    "time"
)

func main(){
    cs := make(chan bool)
    cm := make(chan string)

    go send_message(cm)
    go send_singal(cs)

    // Putting a sleep ensures that we generate message before receiver gets activated.
    // Otherwise, default portion might get executed.
    time.Sleep(time.Second*2)

    // Example of non-blocking message receiver
    select {
        case msg := <-cm:
            fmt.Println("received message", msg)
        default:
            fmt.Println("no message received")
    }

    // Example of non-blocking message or singal receiver
    select {
        case msg := <-cm:
            fmt.Println("received message", msg)
        case sig := <-cs:
            fmt.Println("received signal", sig)
        default:
            fmt.Println("no activity")
    }
}

func send_singal(c chan bool){
    c <- true
}

func send_message(c chan string){
    c <- "Computed result"
}