package main

import (
    "fmt"
    "strconv"
)

func main() {
    // 64 tells how many bits of precision to parse for. What does it mean?
    f, _ := strconv.ParseFloat("123.9", 64)
    fmt.Println(f)

    // The 0 means infer the base from the string. 64 requires that the result fit in 64 bits.
    // I did not get it at all.
    n, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(n)

    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

     _, e := strconv.Atoi("wat")
    fmt.Println(e)
}
