package main

import (
    "time"
    "fmt"
)

func main(){
    ticker1 := time.NewTicker(time.Millisecond * 1000)
    go func() {
       for t := range ticker1.C {
           fmt.Println("Ticked at: ", t)
       }
    }()

    time.Sleep(time.Second * 30)
    ticker1.Stop()
}

