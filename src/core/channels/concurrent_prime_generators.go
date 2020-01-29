package main

import (
    "math/rand"
    "fmt"
)

func main(){
    test_concurrency()
}

func test_concurrency() {
    c1 := make(chan int)
    c2 := make(chan int)
    c3 := make(chan int)
    go prime_producer(100, 10000, c1)
    go prime_producer(100, 10000, c2)
    go prime_producer(100, 10000, c3)

    // select statement picks from set of values. In this case,
    // it is going to pick one frome the list which is having
    // value.
    for i:=0; i<3; {
        select {
        case r1 := <-c1:
            i = i + 1
            fmt.Printf("First one is done, result=%d\n", r1)
        case r2 := <-c2:
            i = i + 1
            fmt.Printf("Second one is done, result=%d\n", r2)
        case r3 := <-c3:
            i = i + 1
            fmt.Printf("Third one is done, result=%d\n", r3)
        }
    }

    fmt.Println("Job completed")
}

func prime_producer(count int, end int, c chan int){
    primes := make([]int, count)
    i := 0
    for j:=0; j<count; {
        number := rand.Intn(end)
        isPrime := check_primness(number)
        if isPrime == true {
            primes[j] = number
            j = j + 1
        }
        i = i + 1
    }

    c <- i
}


func check_primness(num int) bool {
    for div:=2; div<num; div++{
        if num % div == 0 {
            return false
        }
    }
    return true
}
