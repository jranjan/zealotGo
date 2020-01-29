package main

import (
    "time"
    "fmt"
    "math/rand"
)

func main() {
    test_game()
}

type Ball struct {
    hits int
}

func test_game() {
    c := make(chan *Ball)
    ball := new(Ball)
    ball.hits = 0

    go play("Babita", c)
    go play("Bablu..........", c)
    time.Sleep(5*time.Second)
    c <- ball
    time.Sleep(60*time.Second)
    c <- ball
    fmt.Println(ball.hits)
}

func play(name string, c chan *Ball){
    for {
        b := <-c
        b.hits = b.hits + 1
        fmt.Printf("%s served. hits so far=%d\n", name, b.hits)
        wait := time.Duration(rand.Intn(2))
        time.Sleep(wait * time.Millisecond)
        c <- b
    }
}