package main

import (
    "fmt"
    "regexp"
)

func main(){
    s := "hell hi 1234 ok fine 56789 done"
    var pat string
    fmt.Print("enter same pattern:")
    fmt.Scan(&pat)
    cpat,_ := regexp.Compile(pat)
    result := cpat.FindAllString(s, -1)
    fmt.Println("result = ", result)

    n := len(result)
    fmt.Println(n)
}
