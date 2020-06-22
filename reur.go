package main

func fact(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * fact(x-1)
}
func main() {
	fmt.println(fact(10))
}
