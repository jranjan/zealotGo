package main

import (
    "fmt"
    "time"
)

func my_coroutine(msg string){
    for i :=0; i< 10; i = i + 1{
        fmt.Println(msg, ":", i )
        time.Sleep(10)
    }
}

func main(){
    my_coroutine("Direct")
    go my_coroutine("First invoked") // Effectively, you are creating a parallel thread here.
    go func(msg string){
        for i :=0; i< 10; i = i + 1 {
            fmt.Println(msg, ":", i)
            time.Sleep(20)
        }
    }("Hello world!")

    // Let us wait to see all strings printed on our console before
    // we say that we are done.
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}
