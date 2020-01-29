package main

import "fmt"

func main(){
	myarray := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("my array is", myarray)
	avg := avg(myarray)
	fmt.Println("average = ", avg)
}

func avg(numbers []int) int{
	i := 0	
	sum := 0
	for i < len(numbers){
		sum = sum + numbers[i]
		i = i + 1
	}
	
	return sum/len(numbers)
}
