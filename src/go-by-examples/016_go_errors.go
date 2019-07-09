package main

import (
    "errors"
    "fmt"
)

func main() {
    for _, i := range[] int{7, 42, 62}{
        if val, err := f1(i); err != nil {
            fmt.Println("call of f1 failed, passed arg=", i)
        } else {
            fmt.Println("call of f1 worked, result=", val)
        }
    }

    for _, j := range[] int{60, 62, 64}{
        if val, err := f2(j); err != nil {
            fmt.Println("call of f2 failed. Error: ", err)
        } else {
            fmt.Println("Call of f2 worked. result=", val)
        }
    }

    _, e := f2(62)
    if ae, ok := e.(argError); ok {
        fmt.Println(ae.val)
        fmt.Println(ae.message)
    }
}

func f1(arg int) (int, error){
    if (arg == 62){
        return -1, errors.New("Invalid argument, can not work with 62.")
    }
    return arg, nil
}

type argError struct {
    val int
    message string
}

func (e argError) Error() string {
    // You are defining a method here.
    // Your method can take an address or value as an argument. It is up to you.
    // Note that method Error() is language defined.
    // Does it mean that there are two sets of methods: one with pointer and another as it is?
    return fmt.Sprintf("arg=%d, error=%s", e.val, e.message)
}

func f2(arg int)(int, error){
    if (arg==62){
        return -1, argError{arg, "invalid argument"}
    }

    return arg, nil
}

