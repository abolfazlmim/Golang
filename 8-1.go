package main

import "fmt"

func zero(x *int) {
	*x = 50
}

func main() {
	x := new(int)
	zero(x)
	fmt.Println(*x)
}
