package main

import "fmt"

func main() {
	//creating Dynamic arrays and slices
	prices := []float64{}
	prices = append(prices, 90.5, 78)
	fmt.Println(prices)
}

// func main() {
// 	var priceNames [6]string
// 	prices := [4]float64{0.9, 0.2, 11.3, 2.6}
// 	priceNames[5] = "KARIZZZZ"

// 	//slicing an array
// 	featuredPrices := prices[1:3]
// 	fmt.Println(featuredPrices)
// }
