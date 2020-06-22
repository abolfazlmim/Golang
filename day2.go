package main

import "fmt"

var total float64 = 0

func main() {
	var x [5]float64
	for i := 0; i < 5; i++ {
		x[i] = float64(i) + 25
	}
	fmt.Println(x)
	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total)
}
