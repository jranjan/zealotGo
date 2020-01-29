package main

import "fmt"

func main(){
    fmt.Print("Enter a:")
    var a int
    fmt.Scan(&a)
    fmt.Print("Enter a:")
    var b int
    fmt.Scan(&b)
    c := a + b
    fmt.Printf("%d, %d, %d", a, b, c)
}
