package main

import (
	"fmt"
	"example.com/array-practice/product"
)

func main() {
	hobbies := [3]string{"Reading", "Training", "Coding"}
	fmt.Println(hobbies)
	firstElement := hobbies[0]
	fmt.Println("The first element is", firstElement)
	secondAndThird := hobbies[1:]
	fmt.Println("The second and third elements are:", secondAndThird)
	firstSlice := hobbies[0:2]
	secondSlice := hobbies[:2]
	fmt.Println("The first and second slice have the same elements i.e", firstSlice, secondSlice)

	dynArray := []string{"Learn Go", "Create REST APIs with Go"}
	dynArray[1] = "Create REST APIs"
	dynArray = append(dynArray, "Master Go and use it in many of my projects.")
	fmt.Println(dynArray)

	firstProduct, err := product.New("Phone", "phn", 113.65)
	secondProduct, err := product.New("Television", "tv", 123.65)
	if err != nil {
		fmt.Println(err)
	}

	productList := []product.Product{
		*firstProduct, *secondProduct,
	}

	fmt.Println(productList)

	thirdProduct, err := product.New("Fridge", "fr", 123.65)
	productList = append(productList, *thirdProduct)

	fmt.Println(productList)
}
