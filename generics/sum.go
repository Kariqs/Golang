package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)
}

func add[T int | float64](a, b T) T {
	return a + b
}

//In the above function we want the user to input any value and the function should return the type of the input value
