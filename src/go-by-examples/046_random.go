package main

import (
    "fmt"
    "math/rand"
    "time"
)

// The following points to be noted about random numbers:
//      1. rand.x() produces pre-defined sequence
//      2. if we want rand.x() not to be predictable then we should provide a CHANGING seed.
//      3. Do not use rand.x() for secret data. It is safe to use crypto module.
func main(){
    fmt.Println()
    for i := 1; i <= 100; i = i + 1{
        fmt.Printf("%d\t", rand.Intn(100))
        if i % 10 == 0{
            fmt.Println()
        }
    }

    fmt.Println(rand.Float64())

    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    fmt.Print(r1.Intn(100), ",")
    fmt.Print(r1.Intn(100))
    fmt.Println()

    // Illustrate that same source will produce same set of random numbers
    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Print(r2.Intn(100), ",")
    fmt.Print(r2.Intn(100))
    fmt.Println()
    s3 := rand.NewSource(42)
    r3 := rand.New(s3)
    fmt.Print(r3.Intn(100), ",")
    fmt.Print(r3.Intn(100))
}
