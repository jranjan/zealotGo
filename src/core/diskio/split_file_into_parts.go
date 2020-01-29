package main

import (
    "os"
    "fmt"
    "strconv"
)

func main(){
    n := len(os.Args)
    root := "C:\\temp\\"
    if n != 3 {
        fmt.Println("Need at least two inputs")
        os.Exit(1)
    }

    fmt.Println(os.Args[1:])
    parts, err := strconv.ParseInt(os.Args[2], 10, 64);
    if err != nil {
        fmt.Println("Inalid input")
        os.Exit(1)
    }

    fmt.Println(os.Args[1:])
    fmt.Printf("%s will be splitted in %d parts\n", root+os.Args[1], parts)

    fin, err := os.Stat(root+os.Args[1])
    if err != nil{
        panic(err)
    }

    segment_size := fin.Size() / parts
    fmt.Printf("File contains %d bytes, will be splitted in segments of %d bytes\n", fin.Size(),segment_size)
    do_split(root+os.Args[1], root+os.Args[1], segment_size)
}

func do_split(source string, dest string, segment_size int64){
    fin, err := os.Open(source)
    if err != nil{
        panic(err)
    }
    defer fin.Close()

    b := make([]byte, 1)
    count := 1


    fout, err := os.Create(dest + strconv.Itoa(count))
    if err != nil {
        panic(err)
    }

    var dest_bytes int64 = 1
    for {
        n, _ := fin.Read(b)
        if n == 0 {
            break
        }

        if dest_bytes % segment_size == 0 {
            // create a new file, if required
            fout, err = os.Create(dest + strconv.Itoa(count))
            count = count + 1
            if err != nil {
                panic(err)
            }
            defer fout.Close()
        }


        fout.Write(b)
        dest_bytes = dest_bytes + 1
    }
}

