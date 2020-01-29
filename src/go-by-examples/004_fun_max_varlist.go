package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 85, 2}
	fmt.Println("maximum of: ", a)	
	fmt.Println(mymax(a...))
}


func mymax(numbers ...int)(int){
	max := numbers[0]
	for _, curr := range(numbers){
		if curr > max {
			max = curr
		}
	}
	return max
}