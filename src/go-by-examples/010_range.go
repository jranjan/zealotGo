package main

import "fmt"

func main() {
	numbers := [3]int{10, 100, 1000}
	for i, num := range(numbers[:2]){
		fmt.Printf("index=%d, number=%d\n", i, num)
	}

	for i, c := range "golanguage" {
       		fmt.Printf("[%d]=0x%x,%b\n", i, c, c)
    	}
}


