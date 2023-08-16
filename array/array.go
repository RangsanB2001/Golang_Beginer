package main

import "fmt"

var productName [4]string
var price [4]float32

func main() {

	productName[0] = "Iphone"
	productName[1] = "MacBook"
	productName[2] = "ipad"
	productName[3] = "AirPods"

	price := [4]float32{2000, 3000, 40000, 8900}

	fmt.Println(productName)
	fmt.Println(price)
}
