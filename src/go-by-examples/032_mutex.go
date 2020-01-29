package main

import (
    "sync"
    "math/rand"
    "sync/atomic"
    "time"
    "fmt"
)

func main() {
    var state = make(map[int]int)

    // The syntax pinched me. It does not look like a conventional programming
    // language. Why will do use & and {}? It is so because we need a pointer
    // (better to say reference) so that other can use it.
    var mutex = &sync.Mutex{}

    var readOps uint64 = 0
    var writeOps uint64 = 0

    for i := 0; i < 100; i = i + 1{
        go func(){
            total := 0
            for {
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key] // state[] is being protected.
                mutex.Unlock()
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond * 10)
            }
        }()
    }

    for j := 0; j < 10; j = j + 1{
        go func(){
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second)

    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)

    mutex.Lock()
    fmt.Print("state:", state)
    mutex.Unlock()
}
