package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main(){
    fmt.Print("Enter input:")
    scanner := bufio.NewScanner(os.Stdin)
    scanner

    scanner.Scan()
    line := scanner.Text()
    words := strings.Split(line, " ")
    fmt.Println(line)
    fmt.Println(len(words))
}
