package main

import "fmt"

var numberInt, numberInt2 int = 1000, 2000
var msg string = "hello"

func main() {
	numberfloat := 25.4
	fmt.Println(numberInt)
	fmt.Println(numberInt2)
	fmt.Println(numberfloat)
	fmt.Println(msg)

	fmt.Println(numberInt + numberInt2)
	fmt.Println(float64(numberfloat) + float64(numberInt))
}
