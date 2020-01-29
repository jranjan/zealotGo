package main

import (
    "os"
    "fmt"
)

func main(){
    n := len(os.Args)
    root := "C:\\temp\\"
    if n < 3 {
        fmt.Println("Need at least two inputs: one destination and source file")
        os.Exit(1)
    }

    fmt.Println(os.Args[1:])
    for i:=2; i<n; i++{
        fmt.Printf("Concatenating %s to %s...\n", root+os.Args[i], root+os.Args[1])
        do_transfer_to(root+os.Args[i], root+os.Args[1])
        fmt.Printf("[%d] Done\n",i)
    }
    fmt.Println()
}


func do_transfer_to(source string, dest string){
    fin, err := os.Open(source)
    if err != nil{
        panic(err)
    }
    defer fin.Close()

    fout, err := os.OpenFile(dest, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil{
        panic(err)
    }
    defer fout.Close()

    b := make([]byte, 1)
    for {
        n, _ := fin.Read(b)
        if n == 0 {
            break
        }
        fout.Write(b)
    }
}

