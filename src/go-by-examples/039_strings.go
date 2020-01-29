package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println("contains: ", strings.Contains("miniprojects", "es"))
    fmt.Println("compare: ", strings.Compare("miniprojects", "testing"))
    fmt.Println("split: ", strings.Split("a-b-c-d-e", "-"))
    fmt.Println("Replace: ", strings.Replace("foo", "o", "123456", -1))
    fmt.Println("Repeat: ", strings.Repeat("a", 5))
    fmt.Println("Join: ", strings.Join([]string{"miniprojects", "ing"}, "-"))
}
