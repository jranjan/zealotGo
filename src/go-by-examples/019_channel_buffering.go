package main

import (
    "fmt"
    "time"
)

func buffered_sender(tunnel chan string) {
    // Buffered sender will pump as many as it can. Once channel is filled,
    // it will have to wait for channel to be empty so that it can
    // pump again as much as it can. So, behavior of sender is indirectly
    // controlled by receiver.
    for i:= 0; i < 9; i = i + 1   {
        msg := fmt.Sprintln("...Routed gift number is ", i)
        fmt.Println(msg)
        tunnel <- msg
    }
    fmt.Println("Routed all gifts!")
    fmt.Println("Did you see it did not get print unless and until reciever is there, just comment receiver code")
}

func buffered_receiver(tunnel chan string){
    // Hmm!
    // Did you see that function argument? I thought it will be <name type> but it is not so.
    count := 0
    for {
        msg := <- tunnel
        fmt.Printf("Recieved %s\n", msg)
        count = count + 1
        if count % 3 == 0{
            fmt.Println("Tired...sleeping sometime....")
            time.Sleep(30 * time.Second)
            fmt.Println("Let us start work again....")
        }
    }
}

func main() {
    // Create a unbuffered channel. It will block a reciever till someone injects
    // something in channel.
    channel := make(chan string, 3)

    go buffered_sender(channel)
    go buffered_receiver(channel)

    // Let us wait to see all strings printed on our console before
    // we say that we are done.
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}

