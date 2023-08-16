package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("product =", product)
	///add info

	product["Macbook"] = 40000
	product["Macbook Pro"] = 50000
	product["Macbook Air"] = 65000

	fmt.Println("product =", product)
	//delete
	delete(product, "Macbook Pro")
	fmt.Println("product =", product)
	//update
	product["Macbook Air"] = 55000
	fmt.Println("product =", product)

	value1 := product["Macbook"]
	fmt.Println("Value1 =", value1)

	courseName := map[string]string{"101": "Java", "102": "GO", "103": "C#"}
	fmt.Println(courseName)
}
