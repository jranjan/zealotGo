package main

import (
    "fmt"
    "sync"
)

func main (){
    // test_race()
    test_without_race()
}

type Account struct {
    amount int
}

var mutex = &sync.Mutex{}

func test_without_race(){
    a := new(Account)
    a.amount = 0
    go depositor(a)
    go depositor(a)

    // Let us wait to see all strings printed on our console before
    // we say that we are done.
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")

    fmt.Println(a.amount)
}

func test_race(){
    a := new(Account)
    a.amount = 0
    go depositor(a)
    go depositor(a)

    // Let us wait to see all strings printed on our console before
    // we say that we are done.
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")

    fmt.Println(a.amount)
}

func depositor(a *Account){
    for i:=0; i<1000; i++{
        mutex.Lock()
        a.amount = a.amount + 1
        mutex.Unlock()
    }
}