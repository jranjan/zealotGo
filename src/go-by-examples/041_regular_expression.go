package main

import (
    "regexp"
    "fmt"
    "bytes"
)

func main() {
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

    // To use a regular expression, you need to compile. What does compliation do?
    // Once you compile, you can run a variety of functions.
    r, _ := regexp.Compile("p([a-z]+)(ch)")
    fmt.Println(r)

    fmt.Println(r.MatchString("peach"))
    fmt.Println(r.FindString("peach peach punch")) // It found only one of the string
    fmt.Println(r.FindStringIndex("peach punch"))
    fmt.Println(r.FindStringSubmatch("peach punch"))
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))
    fmt.Println(r.FindAllString("peach punch pinch", -1))
    fmt.Println(r.FindAllStringIndex("peach punch pinch", -1))
    fmt.Println(r.FindAllString("peach punch pinch", 2))

    fmt.Println(r.Match([]byte("peach")))

    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))

    // What is difference between regexp.Compile() and regexp.MustCompile?
    // Need to find?
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)
}
