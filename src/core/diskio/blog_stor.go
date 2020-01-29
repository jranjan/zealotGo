package main

import (
    "bufio"
    "os"
)

func main(){
    blog_storage_manager("C:\\temp\\blog.txt")
}

func blog_storage_manager(fname string){
    fout, err := os.Create(fname)
    if err != nil{
        panic(err)
    }
    defer fout.Close()
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan(){
        line := scanner.Text()
        if line == "done" {
            break
        }
        fout.Write([]byte(line+"\n"))
    }
}