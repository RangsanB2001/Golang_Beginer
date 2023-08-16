package main

import "fmt"

var myScore int

func main() {
	fmt.Println("grade calcolator")
	fmt.Scanf("%d", &myScore)
	fmt.Println("score =", myScore)

	if myScore >= 80 {
		fmt.Println("A")
	} else if myScore >= 70 {
		fmt.Println("B")
	} else if myScore >= 60 {
		fmt.Println("C")
	} else if myScore >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
