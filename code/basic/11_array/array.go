package main

import "fmt"

func main() {
	// declaring an array of type int
	var array1 [5]int
	fmt.Println("The values of array1 are:", array1)
	fmt.Println("The length of array1 is:", len(array1))
	fmt.Println()

	// adding values into the array
	fmt.Println("Adding values to array1")
	array1[0] = 29
	array1[1] = 17
	array1[2] = 42
	array1[3] = 13
	array1[4] = 56
	fmt.Println("Now the values of array1 are:", array1)
	fmt.Println()

	// Reading all values from array1 using for-range loop
	fmt.Println("Reading all the values from array1 using for-range loop...")
	for index, value := range array1 {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}
	fmt.Println()

	// declaring an array of type string with shorthand method
	cities := [5]string{"New York", "Seattle", "Mumbai", "Sydney", "Montreal"}
	fmt.Println("The values of cities array are:", cities)
	fmt.Println()

	// Reading specific values
	fmt.Println("Reading values from cities array using index position...")
	fmt.Println("The first city is:", cities[0])
	fmt.Println("The 3rd city is:", cities[2])
	fmt.Println("The last city is:", cities[4])
	fmt.Println("The last city can also be found by:", cities[len(cities)-1])
	fmt.Println()

	// changing values
	fmt.Println("Changing the value at cities[0] to Paris...")
	cities[0] = "Paris"
	fmt.Println("The value at cities[0] is:", cities[0])
}
