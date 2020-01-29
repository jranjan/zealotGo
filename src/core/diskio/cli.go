package main

import (
    "os"
    "fmt"
    "strconv"
)

// Command line inputs are to
func main(){
    n := len(os.Args)
    if n != 3 {
        fmt.Println("Need 2 inputs")
        os.Exit(1)
    }

    x, _ := strconv.Atoi(os.Args[1])
    y, _ := strconv.Atoi(os.Args[2])
    fmt.Printf("%d\t%d\n",x, y)
}

