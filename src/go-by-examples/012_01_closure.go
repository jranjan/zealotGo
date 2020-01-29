package main

import "fmt"

// Closures are more like a a function with set of declared variables whose initilization is
// per instance of function. Use it when you want to avoid lengthy computation and want to
// start from where you left. I mean by rest-and-walk computing!

func main() {
	for i := 1; i <= 3 ; i = i+1 {
		fn_power :=  multiply_recursively(i)
		fmt.Printf("Power of %d are:\n", i)
		for j := 0; j < 10 ; j = j+1 {
			fmt.Printf("%d\t", fn_power())
		}
		fmt.Println("\n")
	}

}

func multiply_recursively(x int) func() int {
	// Think like this.
	// There are two global variables for inner function:
	//	newval - an internal / local variable
	//  x - a passed value
	// Inner function gets access till the life of inner function exists
	newval := 1
	return func() int {
		newval = newval * x
		return  newval
	}
}