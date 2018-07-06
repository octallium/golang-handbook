package main

import "fmt"

func main() {
	d := "Dog"
	c := "Cat"

	// checking the value of variables
	if d == "Dog" {
		fmt.Println("Woff")
	} else {
		fmt.Println("I dont know which animal!")
	}

	// You can also chain if / else conditions
	if c == "monkey" {
		fmt.Println("I am a monkey.")
	} else if c == "Dog" {
		fmt.Println("I am a dog.")
	} else if c == "Cat" {
		fmt.Println("Meoww")
	}
}
