package main

import "fmt"

func main() {
	// Variation 1 - conditional loop
	num1 := 0
	fmt.Println("Starting for loop...")
	for i := 0; i < 11; i++ {
		fmt.Println("Num =", num1)
		num1++
	}
	fmt.Println()
	// Variation 2 - Infinite loop
	num2 := 20
	fmt.Println("Entering infinite loop...")
	for {
		// break condition
		if num2 < 10 {
			break
		}
		fmt.Println("Num =", num2)
		num2--
	}
	fmt.Println()
	// Variation 3 - Optional statements
	num3 := 20
	fmt.Println("Loop with optional statements...")
	for num3 <= 30 {
		fmt.Println("Num =", num3)
		num3++
	}
	fmt.Println()
	// Variation 4 - Boolean operators
	cond := true
	num4 := 30
	fmt.Println("Loops with boolean operator...")
	for cond {
		if num4 >= 40 {
			cond = false
		}
		fmt.Println("Num =", num4)
		num4++
	}
}
