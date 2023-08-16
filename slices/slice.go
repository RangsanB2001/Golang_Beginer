package main

import "fmt"

var courseName []string

func main() {
	courseName = []string{"java", "Python", "GO"}
	fmt.Println(courseName)
	courseName = append(courseName, "C", "C#", "CSS", "JavaScript", "Bootstrap")
	fmt.Println(courseName)

	courseWeb := courseName[4:8]
	fmt.Println(courseWeb)

	courseWeb = courseName[:4]
	fmt.Println(courseWeb)
}
