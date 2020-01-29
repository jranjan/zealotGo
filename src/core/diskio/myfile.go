package main

import "os"

func main(){
    test_file()
}

func test_file(){
    transfer_to("C:\\temp\\bootsqm.dat", "C:\\temp\\bootsqm_copy.dat")
}

func transfer_to(source string, dest string){
    fin, err := os.Open(source)
    if err != nil{
        panic(err)
    }
    defer fin.Close()

    fout, err := os.Create(dest)
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