package greeting

// The above statement declares that this file belong to
// the package "greeting"

// import other packages
import "fmt"

// GoodMorning - Note that the name of our function starts with a capital alphabet,
// capital alphabet denotes that the function is exported or visible outside
// the package.
// Now you can understand why Println() has a capital letter in the beginning !
func GoodMorning() {
	fmt.Println("Good Morning, Octallium!")
	fmt.Println()
}

// goodAfterNoon starts with a lowercase alphabet, means that this function is not
// exported, and can only be accessed within the package.
// Calling goodAfterNoon from another package will result in an error
func goodAfterNoon() {
	fmt.Println("Good Afternoon, Octallium!")
	fmt.Println()
}

// CallGoodAfterNoon has access to goodAfterNoon as it belongs to the same file,
// and hence it can call it.
func CallGoodAfterNoon() {
	goodAfterNoon()
}

// GoodNight accepts a string parameter and prints the result to the screen
func GoodNight(name string) {
	fmt.Println("Good Night, ", name)
	fmt.Println()
}

// Factorial accepts an int parameter and also returns an int,
// it calculates the factorial of the input int and returns
// the factorial.
// At this point ignore how the function works, just concentrate
// on the declaration and syntax.
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
