package mylib

// --------------------------------------------------------------------------------------------------------------------
// Goal:
//
// This sample program demonstrates following:
//
//      1. how to create race conditions in our programs.
//      2. tools to detect race condition
//      3. how to avoid race condition using mutext (there are other alternates as well)
//
// If you run the below program (with race condition) multiple times, you will get different values.
//
// --------------------------------------------------------------------------------------------------------------------


import (
    "fmt"
    "runtime"
    "sync"
)


// --------------------------------------------------------------------------------------------------------------------
//
// Go has race condition tool. You need to do following:
//      1. Build your code with -race option (go build -race demo_race_condition.go)
//      2. Run the program
//
// The sample output:
//
//  c:\Jyoti\Livewire\go\src\brownbag>demo_race_condition.exe
//  ==================
//  WARNING: DATA RACE
//  Read at 0x000000f4ba08 by goroutine 7:
//    main.incCounter()
//      c:/Jyoti/Livewire/go/src/brownbag/demo_race_condition.go:42 +0x81
//
//  Previous write at 0x000000f4ba08 by goroutine 6:
//    main.incCounter()
//       c:/Jyoti/Livewire/go/src/brownbag/demo_race_condition.go:51 +0xb2
//
//  Goroutine 7 (running) created at:
//    main.main()
//      c:/Jyoti/Livewire/go/src/brownbag/demo_race_condition.go:28 +0x90
//
//  Goroutine 6 (finished) created at:
//    main.main()
//      c:/Jyoti/Livewire/go/src/brownbag/demo_race_condition.go:27 +0x6f
// ==================
// Final Counter: 4
// Found 1 data race(s)
//
// --------------------------------------------------------------------------------------------------------------------



var (
    // counter is a variable incremented by all goroutines.
    counter int

    // wg is used to wait for the program to finish.
    mwg sync.WaitGroup

    // mutex is used to define a critical section of code.
    mutex sync.Mutex
)

// main is the entry point for all Go programs.
func DemoRaceCondition(bug bool) {
    // Add a count of two, one for each goroutine.
    mwg.Add(2)

    // Set it to zero before invoking any routine
    counter = 0

    if (bug){
        go buggyIncrementCounter(1)
        go buggyIncrementCounter(2)
    } else {
        go safeIncrementCounter(1)
        go safeIncrementCounter(2)
    }

    // Wait for the goroutines to finish.
    mwg.Wait()
    fmt.Println("Final Counter:", counter)
}

// incCounter increments the package level counter variable.
func buggyIncrementCounter(id int) {
    // Schedule the call to Done to tell main we are done.
    defer mwg.Done()

    for count := 0; count < 2; count++ {

        // Capture the value of Counter.
        value := counter

        // Yield the thread and be placed back in queue.
        runtime.Gosched()

        // Increment our local value of Counter.
        value++

        // Store the value back into Counter.
        counter = value
    }
}



// incCounter increments the package level counter variable.
func safeIncrementCounter(id int) {
    // Schedule the call to Done to tell main we are done.
    defer mwg.Done()

    for count := 0; count < 2; count++ {

        mutex.Lock()
        {
            // Capture the value of Counter.
            value := counter

            // Yield the thread and be placed back in queue.
            runtime.Gosched()

            // Increment our local value of Counter.
            value++

            // Store the value back into Counter.
            counter = value
        }

        // Do not forget to invoke it. Otherwise, you will create a deadlock.
        // Anyhow, Go compiler will raise error at build time.
        // Good, isn't it?
        mutex.Unlock()
    }
}
