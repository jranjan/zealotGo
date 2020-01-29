package main

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

func main() {
    p := point{36, 99}
    fmt.Println(p)
    fmt.Printf("%v\n", p)
    fmt.Printf("%+v\n", p)
    fmt.Printf("%#v\n", p)
    fmt.Printf("%T\n", p)
    fmt.Printf("%t\n", true)
    fmt.Printf("%f\n", 78.9)
    fmt.Printf("%q\n", "\"string\"")
    fmt.Printf("%p\n", &p)
    fmt.Printf("|%6d|%6d|\n", 12, 345)
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
    fmt.Fprintf(os.Stderr, "an %s\n", "error")
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
}

