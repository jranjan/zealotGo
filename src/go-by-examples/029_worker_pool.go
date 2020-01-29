package main

import (
    "fmt"
    "time"
)

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)


    // Start three workers. Think go routint called using 'go <X>' as thread call.
    // As there is no job, all will be waiting for it to arrive.
    for i := 0; i < 3; i = i + 1{
        go square_worker(i, jobs, results)
    }
    go publish_result(results)

    // Push the job for which square needs to be computed
    // We are pushing 25 jobs though it can host 100s of those.
    // And, we will close it as well.
    for j := 0; j < 25; j = j + 1 {
        jobs <- j
    }

    // How will range() behave if we do not close the channel?
    close(jobs)


    // Let us wait to see all strings printed on our console before
    // we say that we are done.
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}

func square_worker(id int, jobs <-chan int, results chan<- int){
    fmt.Printf("Started worker: %d\n", id)
    for job := range jobs {
        number := job
        fmt.Printf("Computing square %d of using worker ID = %d ...\n", number, id)
        time.Sleep(time.Second * 2)
        square := number * number
        fmt.Printf("Square of %d is %d, pushing result using worker ID = %d to another channel\n", number, square, id)
        results <- square
    }

    // Close result channel as we do not expect any more job to be processed.
    // There is race condition here. If some worker thread closes though
    // other wants channel to remain open so that it can push data. It will
    // panic in that case. So, you can close blindly results channel.
    // What next?
    //          close(results)
    // Is it necessary to close the channel? May be not?
    fmt.Printf("Stopped worker: %d \n", id)
}

func publish_result_1(results chan int){
    fmt.Println("Results are...")
    for {
        select {
            case r := <- results:
                fmt.Println(r)
            case <-time.After(time.Second * 20):
                fmt.Printf("All results should have come by now")
                close(results)
                break
        }
    }
}



func publish_result(results chan int){
    fmt.Println("Results are...")
    for {
        r := <- results
        fmt.Println(r)
        break
    }
}