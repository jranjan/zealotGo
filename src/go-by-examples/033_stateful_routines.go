package main

import (
    "time"
    "sync/atomic"
    "math/rand"
    "fmt"
)

type readOp struct {
    key int
    resp chan int
}

type writeOp struct {
    key int
    val int
    resp chan bool
}

func main() {
    // Channel-based approach aligns with Go’s ideas of sharing
    // memory by communicating and having each piece of data owned
    // by exactly 1 go routine.
    var readOps uint64 = 0
    var writeOps uint64 = 0

    // Why did we use '*'?
    reads := make(chan *readOp)
    writes := make(chan *writeOp)

    go manage_map(reads, writes)

    // Create 100 readers.
    for r := 0; r < 100; r++ {
        go func() {
            for {
                read := &readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int)}
                // Make a read request
                reads <- read
                // Waiting for read response. OK!
                <-read.resp
                // Increased what? I am confused!
                // Increase read count. As it shared across multiple go routine,
                // we used atomic operation. Ah, I see! So, we use combination
                // of state management constructs: atomic and channel.
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    for w := 0; w < 10; w++ {
        go func() {
            for {
                write := &writeOp{
                    key: rand.Intn(5),
                    val: rand.Intn(1000),
                    resp: make(chan bool)}
                // Make a write request
                writes <- write
                // Wait for write to complete
                <-write.resp
                // Increase write counter
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second)

    // You can read the readOps and writeOps directly as
    // it is being touched by others.
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}


func manage_map(reads chan *readOp, writes chan *writeOp) {
    var state = make(map[int]int)

    // I do not care who is asking. I just care what is being asked.
    // I loop forever. If there is a read request, I will fetch it
    // and feed to resp channel of read.  Similarly for writes, if
    // there is a write request, I will do fetch the value and
    // update the map state. What is purpose of 'write.resp <- true'?
    // Probably, it is to notify that write is done. Hmm! is it?
    for {
        select {
        case read := <-reads:
            read.resp <- state[read.key]
        case write := <-writes:
            state[write.key] = write.val
            write.resp <- true
        }
    }
}