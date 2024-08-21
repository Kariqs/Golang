package main

import "fmt"

func main() {
	var age int = 32
	fmt.Println("Age:", age)
	getAdultYears(&age)
	fmt.Println("Adult Years", age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
