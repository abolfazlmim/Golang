package main

import "fmt"

func main() {
	a, b := add(1, 2, 3)
	fmt.Println(a, b)
	fmt.Println(add(1, 2, 3, 4, 5, 5, 6))

}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
