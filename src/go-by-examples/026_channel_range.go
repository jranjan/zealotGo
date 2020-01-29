package main

import "fmt"

func main() {
    c := make(chan int, 10)
    for i := 0; i < 10; i = i + 1{
        c <- i
    }
    close(c)

    for number := range(c){
        fmt.Println(number)
    }
}
