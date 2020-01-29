// --------------------------------------------------------------------------------------------------------------------
// Did you notice following:
//      All files under a given folder belong to same package
//      I picked the package name as folder (a typical custom being followed) but not necessary
//
// The story is like this:
//
//      1. You declare a struct
//      2. You declare an interface, called as behaviors often
//      3. You implement methods for a given receiver (a struct)
//
//  The moment your struct (or receiver) implements all interface functions in form of methods, your struct is sort
//  of derived type of interface. You can use interface and struct interchangeably wherever you like.
// --------------------------------------------------------------------------------------------------------------------
package mylib

import (
    "math"
    "fmt"
)

// --------------------------------------------------------------------------------------------------------------------
// Define interface(s)
// --------------------------------------------------------------------------------------------------------------------
type Geo interface {
    Area() float64 // Did you know why I did pick Area() though I like area()?
    Perim() float64
}


// --------------------------------------------------------------------------------------------------------------------
// Define structs (analogy class) implementing interface
// --------------------------------------------------------------------------------------------------------------------
type Rect struct {
    Width float64
    Length float64
}

type Circle struct {
    Radius float64
}


// --------------------------------------------------------------------------------------------------------------------
// Implement interfaces (for rectangle and cube)
// --------------------------------------------------------------------------------------------------------------------
func (r Rect) Area() float64 {
    return r.Width * r.Length
}

func (r Rect) Perim() float64 {
    return r.Width + r.Length
}

func (c Circle) Area() float64 {
    return c.Radius * c.Radius
}

func (c Circle) Perim() float64 {
    return 2 * math.Pi * c.Radius
}


// --------------------------------------------------------------------------------------------------------------------
// Code relying on interfaces
// --------------------------------------------------------------------------------------------------------------------
func Measure(g Geo) (float64, float64){
    a := g.Area()
    p := g.Perim()

    return a, p
}



// --------------------------------------------------------------------------------------------------------------------
// Client code
// --------------------------------------------------------------------------------------------------------------------
func DemoInterface() (error){
    r := Rect{2, 4}
    c := Circle{17}

    a, p := Measure(r)
    fmt.Printf("Rectangle geometry: %#v\n", r)
    fmt.Printf("area: %f, perimeter: %f\n", a, p)

    a, p = Measure(c) // Did you notice the usage difference between := and =?
    fmt.Printf("Circle geometry: %#v", c)
    fmt.Printf("area: %f, perimeter: %f\n", a, p)

    // We do not return any error. I just put here to illustrate the following:
    //  1. error is built-in type of go
    //  2. you can return nil
    return nil
}
