package main

import (
    "fmt"
    "math/rand"
)

func main(){
    test()
}

func test(){
    c := make(chan *Rectangle)
    go producer(11, c)
    go consumer(11, c)

    // Let us wait to see all strings printed on our console before
    // we say that we are done.
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}

type Rectangle struct {
    length int
    width int
    area int
    isSquare bool
}

func (r* Rectangle) Area()  {
    r.area = r.length * r.width
}

func (r* Rectangle) CheckSquare()  {
    if r.length == r.width {
        r.isSquare = true
    }
    r.isSquare = false
}

func create_rect(l int, w int) (*Rectangle){
    r := new(Rectangle)
    r.length = l
    r.width = w
    return r
}

func producer(count int, c chan *Rectangle){
    for i:=0; i<count; i++ {
        r := create_rect(rand.Intn(100), rand.Intn(20)+1)
        r.Area()
        r.CheckSquare()
        c <- r
    }
}

func consumer(count int, c chan *Rectangle){
    for {
        r := <- c
        fmt.Printf("Rectangle[%dx%d], Area=%d, Square=%t\n", r.length, r.width, r.area, r.isSquare)
    }

}