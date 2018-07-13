package main

// The above declaration says that this file belongs to the "main" package

// importing our packages
import (
	// package "fmt" comes in-built as a part of Go standard library
	"fmt"
	// importing our custom package
	"./greeting"
	// the "./" denotes that from the current directory go to the "greeting" directory
)

func main() {
	// Calling the GoodMorning() from greeting package
	fmt.Println("Calling function GoodMorning() from package 'greeting'...")
	greeting.GoodMorning()
	// Calling the goodAfterNoon() from greeting package
	fmt.Println("Calling function goodAfterNoon() from package 'greeting'...")
	// as goodAfterNoon is not directly accessable, we call CallGoodAfterNoon()
	// which in turn calls goodAfterNoon()
	greeting.CallGoodAfterNoon()
	// Calling the GoodNight() from greeting package
	fmt.Println("Calling function GoodNight() from package 'greeting'...")
	greeting.GoodNight("Gopher!")
	// Calling the Factorial function with passing an int
	num := 10
	// As Factorial() returns an int, we save it in a new variable "result"
	result := greeting.Factorial(num)
	// Printing the result on screen, note we are using Printf()
	fmt.Printf("Factorial of %d is: %d.\n", num, result)
}
