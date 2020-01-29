package main

import (
    "os"
    "fmt"
)

func main() {
    f := createFile("C:\\defer.txt")
    defer closeFile(f)  // It gives freedom to call the function as late as possible.
                        // Not sure, what other benefits it bring?
    writeFile(f)

    // Do not expect other thread to wait for completion.
    // So, calls are in context of parent thread.
    go test_routine()
    fmt.Println(".......................................")
}


func createFile(p string) *os.File{
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File){
    fmt.Println("writing")
    fmt.Fprintln(f, "data")
}

func closeFile(f *os.File){
    fmt.Println("closing")
    f.Close()
}

func test_routine(){
    for {
        fmt.Println("Test")
    }
}
