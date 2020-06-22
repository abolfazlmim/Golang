package main

import "fmt"

func main() {
	var x [10]int
	var total float64 = 0
	x[4] = 100
	fmt.Println(x)
	for i := 0; i < 10; i++ {
		x[i] = i + 250
		total += float64(x[i])
	}
	fmt.Println(x)
	fmt.Println(total)
}
