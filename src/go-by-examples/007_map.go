package main

import "fmt"

func main() {
	mymap := make(map[string]int)

	mymap["Jyoti"] = 1000
	mymap["Unmesh"] = 10000
	mymap["Saktiman"] = 100000
	mymap["Pradeep"] = 1000000

	fmt.Println(mymap)
	delete(mymap, "Saktiman")
	fmt.Println(mymap)
	
	jv, prs := mymap["Jyoti"]
	fmt.Println("prs:", prs, "Jyoti's value:", prs, jv)
}
