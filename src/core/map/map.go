package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

// Problem statements:
// Write a program to remove duplicate words from a string
// Write a program to find frequency occurence of words from a string
// Write a program to read n strings and search a word in list of string


func main(){
    // remove_duplicates()
    // count_frequency()
    display_multiple_lines
}

func display_multiple_lines(){
    lines := read_multiple_lines("Enter your input")
    fmt.Println(lines)
}

func count_frequency(){
    line := read_string("Enter your string")
    word_map := create_word_map(line)
    key := read_string("Enter key string")
    count, flag := word_map[key]
    if flag == false{
        fmt.Printf("%s does not exist\n")
        return
    }

    fmt.Printf("The word [%s] exists %d times.", key, count)
}

func remove_duplicates(){
    line := read_string("Enter your string")
    word_map := create_word_map(line)
    fmt.Println(word_map)
    words := strings.Split(line, " ")
    mystring := []string{}
    for _, w := range words {
        if word_map[w] == 1{
            mystring = append(mystring, w)
        }
    }
    strings.Join(mystring, " ")
    fmt.Println(mystring)
}

func create_word_map(s string) map[string]int {
    words := strings.Split(s, " ")
    m := map[string]int{}
    for _, w := range words {
        _, flag := m[w]
        if flag == false {
            m[w] = 1
        } else {
            m[w] = m[w] + 1
        }
    }
    return m
}

func read_string(msg string) (string){
    fmt.Printf("Enter %s:\n", msg)
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    line := scanner.Text()
    return line
}


func read_multiple_lines(msg string) []string{
    fmt.Printf("Enter %s:\n", msg)
    lines := []string{}
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan(){
        line := scanner.Text()
        if line == "end" { break }
        lines = append(lines, line)
    }

    return lines
}