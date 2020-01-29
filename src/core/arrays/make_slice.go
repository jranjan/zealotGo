package main

import "fmt"

func main(){
    make_v1()
    //m := allocate_2D_matrix(3, 6)
    //fmt.Println(m)
}

func make_v1(){
    s := make([]int, 2, 5)
    fmt.Println(s)
    for i:=0; i<10; i++{
        if i >= len(s){
            s = append(s, 10*i)
        }
        s[i] = 10 * i
        fmt.Printf("Length=%d, Capacity=%d, Value=%d\n", len(s), cap(s), s[i])
    }
}

func allocate_2D_matrix(row int, col int) [][]int{
    m := make([][]int, row)
    for i:=0; i<row; i++{
        m[i] = make([]int, col)
    }
    return m
}