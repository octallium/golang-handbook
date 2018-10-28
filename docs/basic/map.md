# Understanding Data Structure of Type {++ Map++}.

<hr>

## Objective

Understanding {++ Map's++} in Go.

![Maps](../images/maps-octallium.jpg)

Imagine yourself creating a team of Avengers and need to add contact details of all your Super Hero's, you essentially would require a phone number be associated with only one Super Hero, this can be easily done using 'Maps' in Go.

It is also known as 'Hash Table' in other programming languages and these offer faster lookups, adds and deletes.

## Structure

Navigate to our code folder

    code/basic/

For our program create a new folder '13_map'

    code/basic/13_map/

And lets create a file 'maps.go' in it, finally the structure would look like this:

    code/basic/13_map/maps.go

## Declaration

Syntax

Declaration & initialization method

    mapName := map[KeyType]ValueType

With built-in function make()

    map := make(map[KeyType]ValueType)

## Code

We will write the code in 2 parts:

1.)

???+example "part-1 maps.go"

    ``` go
    1 package main
    2
    3 import "fmt"
    4
    5 func main() {
    6	// Let's make contact list with names & phone numbers
    7	// Initialize a map 'contactList'
    8	contactList := make(map[string]int)
    9	// Adding values to map
    10	contactList["Iron Man"] = 878111222
    11	contactList["Thor"] = 121131141
    12  contactList["Batman"] = 483910138
    13  contactList["Spider Man"] = 478282929
    14
    15    // print out the map
    16    for key, value := range contactList {
    17        fmt.Printf("Key = %s\tValue = %d\n", key, value)
    18    }
    19    fmt.Println("")
    20    // Iron man decides to change his number
    21    contactList["Iron Man"] = 333333333
    22    fmt.Printf("The new contact no of Iron Man is: %d\n", contactList["Iron Man"])
    23    fmt.Println("")
    24    // Now Hulk wants to join the team
    25    contactList["Hulk"] = 911831925
    26
    27    // New team is
    28    for key, value := range contactList {
    29        fmt.Printf("Key = %s\tValue = %d\n", key, value)
    30    }
    31
    ```

_Review_

On line 8 we declare a 'map' and initialize using 'make'

    contactList := make(map[string]int)

We create a map with 'key' of type 'string' and 'value' of type 'int', we want to map the name of the super hero along with their contact number.

We can add values to map as done from line 10 to 14.

    contactList["Iron Man"] = 878111222
    contactList["Thor"] = 121131141
    contactList["Batman"] = 483910138
    contactList["Spider Man"] = 478282929

Map can also be declared and initialized using the following syntax

    contactList := map[string]int {
    	"Iron Man": 878111222,
    	"Thor": 121131141,
    	"Batman": 483910138,
    	"Spider Man": 478282929,
    }

On line 16 we print out the map. We can also change the values of the key, on line 21 we change the number of Iron Man.

    contactList["Iron Man"] = 333333333

We can also add new super hero on line 25

    contactList["Hulk"] = 911831925

2.)

???+example "part-2 maps.go"

    ```go
    32	/*
    33        Operators in maps
    34    */
    35    // Length of our super hero team
    36    fmt.Println("Length: ", len(contactList))
    37
    38    // Spider man is not performing well and we need to delete him
    39    delete(contactList, "Spider Man")
    40    fmt.Println("Length: ", len(contactList))
    41
    42    // Checking if Spider Man is deleted or not
    43    _, ok := contactList["Spider Man"]
    44    if !ok {
    45        fmt.Println("Spider Man is deleted")
    46    }
    47 }
    ```

_Review_

On line 36 we check the length of map.

    fmt.Println("Length: ", len(contactList))

We can delete values from our map using 'delete' as on line 39.

    delete(contactList, "Spider Man")

On line 43 we check if the key is present in the map or not, it returns a boolean value, 'true' if the key is present and 'false' if not found.

    _, ok := contactList["Spider Man"]

## Full Code

??? example "slice.go"

    ``` go
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

        // Siper man is not performing well and we need to delete him
        delete(contactList, "Spider Man")
        fmt.Println("Length: ", len(contactList))

        // Checking if Siper Man is deleted or not
        _, ok := contactList["Spider Man"]
        if !ok {
            fmt.Println("Spider Man is deleted")
        }
    }

    ```

## Running your code

Open your terminal and navigate to our folder

    code/basic/13_map/

Once in the folder type the following command

    go run maps.go

## Output

If there are no errors, you should get the output as:

???+ success "Output"

        Key = Iron Man  Value = 878111222
        Key = Thor      Value = 121131141
        Key = Batman    Value = 483910138
        Key = Spider Man        Value = 478282929

        The new contact no of Iron Man is: 333333333

        Key = Iron Man  Value = 333333333
        Key = Thor      Value = 121131141
        Key = Batman    Value = 483910138
        Key = Spider Man        Value = 478282929
        Key = Hulk      Value = 911831925
        Length:  5
        Length:  4
        Spider Man is deleted

If for some reason your code isn't working, checkout the github repo or playground.

## Github

[Github Repo](https://github.com/octallium/golang-handbook/tree/master/code)

## Golang Playground

[Golang Playground](https://play.golang.org/p/UP11dLv3kFR)

## Next

In the next chapter we will study {++ Struct++}.

## Please Donate ❤️

All the work is provided free of cost and completely open source, but it needs your support and love to keep the activity sustainable.

Any support is genuinely appreciated, you can help by sending a small donation by clicking the below link:

[![PayPal](../images/paypal-logo.png)](https://www.paypal.me/octallium)
