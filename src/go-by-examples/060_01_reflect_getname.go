package main

import (
    "fmt"
    "reflect"
)


// Here's a basic interface for geometric shapes.
type geometry interface {
    area() float64
    perim() float64
}

// For our example we'll implement this interface on
// `rect` and `circle` types.
type rect struct {
    width, height float64
}

// To implement an interface in Go, we just need to
// implement all the methods in the interface. Here we
// implement `geometry` on `rect`s.
func (r rect) area() float64 {
    return r.width * r.height
}

func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func getName(r *rect) string {
    s := reflect.TypeOf(r)
    return fmt.Sprintf("%s", s)
}

func main(){
	 r := rect{width: 3, height: 4}
     fmt.Println(r.height)
     fmt.Println(getName(&r))
}
