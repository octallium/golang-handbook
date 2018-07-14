package main

import "fmt"

func main() {
	// declaring a nil slice
	var slice1 []int
	fmt.Println("slice1:", slice1)
	fmt.Println("The length of slice1 is:", len(slice1))
	fmt.Println("The capacity of slice1 is:", cap(slice1))
	fmt.Println()
	// declaring a slice with initialization
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice2:", slice2)
	fmt.Println("The length of slice2 is:", len(slice2))
	fmt.Println("The capacity of slice2 is:", cap(slice2))
	fmt.Println()

	// declaring a slice of length 5 with make
	slice3 := make([]string, 5)
	fmt.Println("slice3:", slice3)
	fmt.Println("The length of slice3 is:", len(slice3))
	fmt.Println("The capacity of slice3 is:", cap(slice3))
	fmt.Println()

	// declaring a slice of length 5 and capacity 10 with make
	slice4 := make([]int, 5, 10)
	fmt.Println("slice4:", slice4)
	fmt.Println("The length of slice4 is:", len(slice4))
	fmt.Println("The capacity of slice4 is:", cap(slice4))
	fmt.Println()

	// inserting values, note i < 6 will give an error as we have
	// set the length to 5
	for i := 0; i < 5; i++ {
		slice4[i] = i
	}
	fmt.Println("slice4:", slice4)
	fmt.Println()

	// increasing the length of slice
	fmt.Println("Increasing the length of slice...")

	// slice4 = slice4[:11] will give an error as capacity is 10
	slice4 = slice4[:10]
	fmt.Println("The length of slice4 is:", len(slice4))
	fmt.Println("The capacity of slice4 is:", cap(slice4))
	for i := 5; i < 10; i++ {
		slice4[i] = i
	}
	fmt.Println()

	// printing slice4
	fmt.Println("slice4:", slice4)
	fmt.Println()

	// creating a new slice
	slice5 := slice4[2:8]
	fmt.Println("slice5:", slice5)
	fmt.Println()

	// two-D slice
	twoD := [][]int{{3, 4}, {1, 5}, {9, 2}, {7, 8}}
	fmt.Println("Print out values of twoD slice...")
	for i, subSlice := range twoD {
		fmt.Printf("At index: %d of twoD, Value: %v\n", i, subSlice)
		for index, value := range subSlice {
			fmt.Printf("Index: %d Value: %d\n", index, value)
		}
		fmt.Println()
	}
}
