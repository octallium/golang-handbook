package main

import "fmt"

func main() {
	// Declaring variable "name" of type string
	var name string
	name = "Octallium"
	// Printing out the value of "name"
	fmt.Println("The value of 'name' variable is:", name)

	// Declaring variable "age" of type int
	var age int
	age = 9
	// Printing out the value of "age"
	fmt.Println("The value of 'age' variable is:", age)

	// Declaring constant "SECRET_KEY" of type string
	const SECRET_KEY string = "abc-123-xyz-098"
	// Printing out the value of "SECRET_KEY"
	fmt.Println("The value of 'SECRET_KEY' constant is:", SECRET_KEY)

	// Values of variables can be changed
	name = "Golang Handbook"
	fmt.Println("Now the value of 'name' variable is:", name)

	age = 13
	fmt.Println("Now the value of 'age' variable is:", age)

	// SECRET_KEY = "que-472-ert-383" will throw an error as it is a constant
}
