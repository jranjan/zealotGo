package main

import (
    "sort"
    "fmt"
    "os"
)

type byLength []string

func (s byLength) Len() int{
    return len(s)
}

func (s byLength) Swap(i int, j int)  {
    s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i int, j int) bool{
    return  len(s[i]) < len(s[j])
}


// If you want custom sort then you need to do following:
//      1. Define a new data type
//      2. Define sort interface
// Did you notice the sort.Sort() argument?
func main() {
    var country = []string{ "India", "US", "Pakistan"}
    sort.Sort(byLength(country))
    fmt.Println(country)
}


