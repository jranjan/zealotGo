package main

import "fmt"

func main ()  {
    test_decorator()
}

func test_decorator(){
    s := adder(2,3)
    fmt.Println(s)
    my_adder := do_addon(adder, 20)
    fmt.Printf("%d\n", my_adder(2,4))
}

// the function which is going to be decorated. Let us call it decoratee.
func adder(a int, b int) int {
    return a + b
}

// The function which will take decoratee. And apply decorator on to of it.
//
// do_addon(decoratee, addon int)
//
// Truth is that this function does not need nested functon as it does not
// do apply the same logic in old context. However, it can be annotated if there has
// to be context to be applied. Beawre it is not done here!
//
// If it confuse you further then just go back to decorator design pattern
func do_addon(f func(int, int) int, addon int) func(int, int) int {
    decor := func(x int, y int) int {
        s := f(x, y) + addon
        return s
    }
    return decor
}