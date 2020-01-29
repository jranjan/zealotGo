package main

import (
    "fmt"
)

func sender(tunnel chan string) {
    for i:= 0; i < 10; i = i + 1   {
        msg := fmt.Sprintln("Routed gift number is ", i)
        tunnel <- msg
    }
    fmt.Println("Routed all gifts!")
    fmt.Println("Did you see it did not get print unless and until reciever is there, just comment receiver code")
}

func receiver(tunnel chan string){
    // Hmm!
    // Did you see that function argument? I thought it will be <name type> but it is not so.
    for i:= 0; i < 10; i = i + 1   {
        msg := <- tunnel
        fmt.Printf(msg)
    }
}

func main() {
    // Create a unbuffered channel. It will block a reciever till someone injects
    // something in channel.
    channel := make(chan string)

    go sender(channel)
    go receiver(channel)

    // Let us wait to see all strings printed on our console before
    // we say that we are done.
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}
