package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("Datasets.csv")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line %d %s", count, line)

		count++
	}
}
