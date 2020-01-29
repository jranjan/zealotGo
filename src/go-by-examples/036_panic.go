package main

import "os"

func main() {
    panic("Error occured")

    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}
