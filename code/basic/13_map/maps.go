package main

import "fmt"

func main() {
	// Let's make contact list with names & phone numbers
	// Initialize a map 'contactList'
	contactList := make(map[string]int)
	// Adding values to map
	contactList["Iron Man"] = 878111222
	contactList["Thor"] = 121131141
	contactList["Batman"] = 483910138
	contactList["Spider Man"] = 478282929

	// print out the map
	for key, value := range contactList {
		fmt.Printf("Key = %s\tValue = %d\n", key, value)
	}
	fmt.Println("")
	// Iron man decides to change his number
	contactList["Iron Man"] = 333333333
	fmt.Printf("The new contact no of Iron Man is: %d\n", contactList["Iron Man"])
	fmt.Println("")
	// Now Hulk wants to join the team
	contactList["Hulk"] = 911831925

	// New team is
	for key, value := range contactList {
		fmt.Printf("Key = %s\tValue = %d\n", key, value)
	}

	/*
		Operators in maps
	*/
	// Length of our super hero team
	fmt.Println("Length: ", len(contactList))

	// Spider man is not performing well and we need to delete him
	delete(contactList, "Spider Man")
	fmt.Println("Length: ", len(contactList))

	// Checking if Spider Man is deleted or not
	_, ok := contactList["Spider Man"]
	if ok == false {
		fmt.Println("Spider Man is deleted")
	}
}
