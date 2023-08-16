package main

import "fmt"

func zerovalue(ivalue int) {
	ivalue = 0
}
func zeropointer(ipiont *int) {
	*ipiont = 0
}
func main() {
	i := 1
	fmt.Println("i =", i)
	zerovalue(i)
	fmt.Println("i zzerovalue =", i)
	zeropointer(&i)
	fmt.Println("i zzerovalue =", i)
	fmt.Println("i zzerovalue =", &i)
}
