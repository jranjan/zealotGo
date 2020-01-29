package main

import "fmt"

func main(){
	var a, c int
	a, _ , c = vals()
	fmt.Printf("%d, -, %d", a, c)
}


func vals() (int, int, int){
	return 10, 100, 1000
}