package main

import (
    "sort"
    "fmt"
)

func main() {
    var message = []string{"A", "B", "C", "a", "b", "c"}
    fmt.Println(message)
    sort.Strings(message)
    fmt.Println("After sorting: ", message)


    var nos = []int{2, 2, 7, -8, 9, 0}
    var sortedNos = []int{1, 2, 3}

    fmt.Println(nos)
    fmt.Println("Array is:", sortedNos, "sort check result:", sort.IntsAreSorted(nos))
    sort.Ints(nos)
    fmt.Println("After sorting: ", nos)

    fmt.Println("Array is:", sortedNos, "sort check result:", sort.IntsAreSorted(sortedNos))

}
