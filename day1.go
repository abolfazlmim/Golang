package main

import "fmt"

func main() {
	var x [10]int
	x[4] = 100
	fmt.Println(x)
	for i := 0; i < 10; i++ {
		x[i] = i + 250
	}
	fmt.Println(x)
}
