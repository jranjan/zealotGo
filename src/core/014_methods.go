package main

import "fmt"

type cube struct {
    length int
    width int
    height int
}

func (c cube) volume() int {
    return (c.length * c.width * c.height)
}

func (c cube) perimeter() int {
    return 4 * (c.length + c.width + c.height)
}

func (c *cube) surfacearea() int {
    return 2 * (c.length*c.width + c.length*c.height + c.width*c.height)
}

func main() {
    mycube := cube{6, 4, 10}
    fmt.Println("my cube attributes are:", mycube)
    fmt.Println("volume = ", mycube.volume())
    fmt.Println("surface area = ", mycube.surfacearea())
    fmt.Println("perimter = ", mycube.perimeter())
}

