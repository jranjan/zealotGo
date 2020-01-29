package main

import (
    "crypto/sha1"
    "fmt"
    "crypto/md5"
)

func main() {
    fmt.Println("Hash...")
    s := "Hello wrold, India is a great country!"
    fmt.Printf("%x\n", []byte(s))

    // Pattern for computing hash:
    //      1. create hash
    //      2. write bytes
    //      3. compute hash: sha1.Sum([]byte{}). Optionally, you can append hash to extra string
    h := sha1.New()
    var bs []byte
    for i := 0; i < 1; i = i + 1{
        // keep writing to hash stream
        h.Write([]byte(s))

        // h.Sum([]byte(s)) will just append has after s
        bs = h.Sum(nil)

        fmt.Printf("%x\n", bs)
    }

    fmt.Println("MD5...")
    do_same_with_md5()
}


func do_same_with_md5(){
    s := "Hello wrold, India is a great country!"
    fmt.Printf("%x\n", []byte(s))

    h := md5.New()
    var bs []byte
    h.Write([]byte(s))
    bs = h.Sum(nil)
    fmt.Printf("%x\n", bs)
}