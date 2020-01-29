package main

import (
    "fmt"
    "strings"
)

func main() {
    // Go doesn't support collections.
    var strs = []string{"peach", "apple", "pear", "plum"}
    fmt.Println(strs)
    fmt.Println(Index(strs, "plum"))
    fmt.Println(Include(strs, "apples"))
    fmt.Println(Any(strs, has_prefix_p))
    fmt.Println(Filter(strs, has_prefix_p))
    fmt.Println(MapStrings(strs, upper))
}


// Is string 't' in provided list of 'vs' string?
func Index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

// Boolean version of Index()
func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}

// Returns true if any of the string in list satisfies predicate f()
func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs{
        if f(v){
            return true
        }
    }

    return false
}

// Returns a new slice applying predicate f() to list
func Filter(vs []string, f func(string) bool) []string{
    vsf := make([]string, 0)
    for _, v := range(vs){
        if f(v){
            vsf = append(vsf, v)
        }
    }

    return vsf
}

// Apply predicate f() and return result along with original data
func MapStrings(vs []string, f func(string) string) map[string]string {
    vsm := make(map[string]string)
    for _, v := range vs {
        vsm[v] = f(v)
    }
    return vsm
}

func has_prefix_p(s string) bool{
    return strings.HasPrefix(s, "p")
}

func upper(s string) string{
    return strings.ToUpper(s)
}

