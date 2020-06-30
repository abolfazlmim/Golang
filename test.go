package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	var sss string
	var ii uint64
	var dd float64
	var text string
	// Read and save an integer, double, and String to your variables.
	for n := 1; n <= 3 && scanner.Scan(); n++ {
		switch n {
		case 1:
			sss = scanner.Text()
			ii, _ = strconv.ParseUint(sss, 0, 64)
		case 2:
			sss = scanner.Text()
			dd, _ = strconv.ParseFloat(sss, 64)
		case 3:
			text = scanner.Text()
		}
	}
	// Print the sum of both integer variables on a new line.
	fmt.Println(i + ii)
	fmt.Printf("%.1f\n", d+dd)
	fmt.Println(s + text)

	// Print the sum of the double variables on a new line.

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.

}
