package main

import (
    "sync/atomic"
    "time"
    "fmt"
)


// Not sure how it does illustrate atomic counters
func main() {
    var ops uint64 = 0

    for i := 0; i < 100; i = i + 1 {
        go func() {
           for {
               atomic.AddUint64(&ops, 1)
               time.Sleep(time.Millisecond * 2)
           }
        }()
    }

    time.Sleep(time.Second * 10)
    opsFinal := atomic.LoadUint64(&ops)
    fmt.Println(opsFinal)
}
