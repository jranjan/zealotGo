package main

import (
    "fmt"
    "reflect"
)

func main(){
    // test_point_move()
    test_interface_type()
}

func test_interface_type(){
    var g interface{}
    fmt.Println("type=", reflect.TypeOf(g), "valuue=", g)
    g = 10
    fmt.Println("type=", reflect.TypeOf(g), "valuue=", g)
    fmt.Println("type=", reflect.TypeOf(v))
    switch g.(type){
        //
    }
}

func test_point_move(){
    p := Point{2,3}
    fmt.Println(p)
    p.MovevVer1(10,20)
    fmt.Println(p)
    p.MovevVer2(10,20)
    fmt.Println(p)
}

type Point struct {
    x int
    y int
}

func (p Point) MovevVer1(x int, y int) {
    p.x = x
    p.y = y
}


func (p *Point) MovevVer2(x int, y int) {
    p.x = x
    p.y = y
}