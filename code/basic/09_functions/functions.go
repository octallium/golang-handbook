package main

import "fmt"

// sayHello prints "Hello, Octy!"
func sayHello() {
	fmt.Println("In func sayHello()...")
	fmt.Println("Hello, Octy!")
	fmt.Println()
}

// function with passing values
func sayHelloTo(name string) {
	fmt.Println("In func sayHelloTo()...")
	fmt.Println("Hello,", name)
	fmt.Println()
}

// function with return values
func printNums(n int) error {
	fmt.Println("In func printNums()...")
	var err error
	for i := 0; i <= n; i++ {
		_, err = fmt.Println(i)
	}
	return err
}

// main function
func main() {
	fmt.Println("In func main()")
	fmt.Println("Now calling func sayHello()...")
	fmt.Println()
	// calling function sayHello()
	sayHello()
	// calling a function with passing values
	name := "Gopher"
	sayHelloTo(name)
	// calling a function with return parameters
	err := printNums(10)
	if err == nil {
		fmt.Println("There are no errors!")
	} else {
		fmt.Println("Error with printing:", err)
	}
}
