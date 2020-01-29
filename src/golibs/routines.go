// This sample program demonstrates how the goroutine scheduler
// will time slice goroutines on a single thread.
package mylib


// --------------------------------------------------------------------------------------------------------------------
// Goal:
//
// This program demonstrates concurrency support in Go language. We are running two set of routines to generate
// a set of prime number prefixed by specific strings (user feeded). As these two routine execute concurrently
// (not in parallel), the generated numbers are intermixed. Primarily, it demonstrates following:
//
//      1) Concurrency using go
//      2) Touch skin for number of logical processor to be used to run concurrent routines
//
// --------------------------------------------------------------------------------------------------------------------

import (
     "fmt"
     "runtime"
     "sync"
     "time"
)

// wg is used to wait for the program to finish.
var rwg sync.WaitGroup

// main is the entry point for all Go programs.
func DemoRoutines() {
    // Allocate 1 logical processors for the scheduler to use.
    runtime.GOMAXPROCS(1)

     // Add a count of two, one for each goroutine.
     rwg.Add(2)

     // Create two go routines.
     fmt.Println("Create Goroutines")

     go printPrime("Jyoti's prime number collection:")
     go printPrime("Vishnu's prime number collection:")

     // Wait for the goroutines to finish.
     fmt.Println("Waiting To Finish")
     rwg.Wait()

     fmt.Println("Terminating Program")
}

// printPrime displays prime numbers for the first set of  numbers.
func printPrime(prefix string) {
     // Schedule the call to Done to tell main we are done.
     defer rwg.Done()

     for outer := 2; outer < 25; outer++ {
         primeFlag := true
         for inner := 2; inner < outer; inner++ {
             if outer%inner == 0 {
                 time.Sleep(time.Millisecond * 200)
                 primeFlag = false
             }
         }
         if (primeFlag) {
             fmt.Printf("%s:%d\n", prefix, outer)
         }
     }
     fmt.Println("Completed", prefix)
}