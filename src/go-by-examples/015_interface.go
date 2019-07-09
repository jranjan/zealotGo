package main

import (
    "fmt"
    "math"
)

type geometry interface {
    area() float64
    perimeter() float64
}

type circle struct {
    radius float64
}

type rect struct {
    length float64
    width float64
}

func (r rect) area() float64 {
    return r.length * r.width
}

func (r rect) perimeter() float64 {
    return 2 * (r.length + r.width)
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
    return 2* math.Pi * c.radius
}

func measure(g geometry){
    fmt.Println("Geometry: ", g)
    fmt.Println("Area: ", g.area())
    fmt.Println("Perimeter: ", g.perimeter())
}


func main() {
    r := rect{10,5}
    c := circle{6}
    measure(r)
    measure(c)
}
