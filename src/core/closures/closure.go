package main

import (
    "fmt"
)

func main(){
    // test_enclosure()
    get_primes()
}

func test_enclosure(){
    fn := next()
    for i:=0; i <10; i++ {
        fmt.Println(fn())
    }
}


func get_primes(){
    fn := get_prime_set()
    for i:=0; i <10000; i++ {
        fmt.Println(fn())
    }
}

func next() func() int {
    i := 0
    inner := func() int {
        i = i + 1
        return i
    }
    return inner
}


func get_prime_set() func() int {
    current_prime := 1
    prime_generator := func() int {
        start := current_prime + 1
        for {
            if check_primness(start) == true {
              current_prime = start
              return current_prime
            }
                start = start + 1
        }
    }

    return prime_generator
}

func check_primness(num int) bool {
    for div:=2; div<num; div++{
        if num % div == 0 {
            return false
        }
    }

    return true
}
