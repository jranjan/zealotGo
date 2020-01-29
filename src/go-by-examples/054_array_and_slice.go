package main

import "fmt"

func main()  {
    da := []int{10,21,29}
    fmt.Printf("%d\n", len(da))
    fmt.Println(da)
    da = append([]int{100,6,7,9,10,11,12}, da...)
    fmt.Println(".......................................")
    fmt.Printf("%d\n", len(da))
    fmt.Println(da)

    s1 := append(da[:2], da[2+1:]...)
    fmt.Println(".......................................")
    fmt.Printf("%d\n", len(s1))
    fmt.Println(s1)
}