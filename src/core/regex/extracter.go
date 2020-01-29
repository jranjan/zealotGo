package main

import (
    "bufio"
    "os"
    "fmt"
    "regexp"
)

func main(){
    test_vowel_words()
}

func test_vowel_words(){
    lines := []string{}
    fmt.Println("Enter inputs:")
    scanner := bufio.NewScanner(os.Stdin)
    for {
        scanner.Scan()
        line := scanner.Text()
        if line == "end" { break }
        lines = append(lines, line)
    }

    for _, line := range lines {
        fmt.Printf("Extracting vowels from line [%s]\n", line)
        extra_vowel_words(line)
    }
}

func extra_vowel_words(s string){
    var pattern string = "\\b[aeiou]\\w*\\b"
    cpat,_ := regexp.Compile(pattern)
    result := cpat.FindAllString(s, -1)
    fmt.Println("result = ", result)
    n := len(result)
    fmt.Println(n)
}
