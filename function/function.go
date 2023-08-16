package main

import "fmt"

func hello() {
	fmt.Println("Hello Pound")
}
func plus(value1 int, value2 int) int {
	return value1 + value2
}
func plus3Value(value1, value2, value3 int) int {
	return value1 + value2 + value3
}
func main() {
	hello()
	result := plus(3, 4)
	fmt.Println("result =", result)

	result2 := plus3Value(4, 4, 3)
	fmt.Println("result =", result2)
}
