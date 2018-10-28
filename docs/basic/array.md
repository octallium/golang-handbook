# Understanding Data Structure of Type {++ Array++}.

<hr>

## Objective

Understanding {++ Array's++} in Go.

![Array](../images/ice-cream-array-octallium.png)

Imagine walking into an ice-cream store, you see all different flavour displayed nicely side-by-side, tempting isn't it ? You can easily see all the flavours and make your choice easily. This display of various ice-cream flavours is nothing by an "array" with the common item type of "ice-cream" !

Go offers a couple of in-build data structures and one of them is an "array", you can imagine array as a collection of similar items, i.e you can group together common items.

The advantage of array is that it gives you easy access to data, also data can be organized properly by using arrays, however arrays are fixed in size, to overcome this Go offers "slice" which we will cover in the next chapter.

## Structure

Navigate to our code folder

    code/basic/

For our program create a new folder '11_array'

    code/basic/11_array/

And lets create a file 'array.go' in it, finally the structure would look like this:

    code/basic/11_array/array.go

## Declaration

Syntax

    var arrayName[size] type

Or declaration & initialization method

    arrayName := [size]type{value1, value2, ...}

## Code

We will write the code in 2 parts:

1.)

???+example "Part-1 array.go"

    ```go
        1 package main
        2
        3 import "fmt"
        4
        5 func main() {
        6 // declaring an array of type int
        7 var array1 [5]int
        8 fmt.Println("The values of array1 are:", array1)
        9 fmt.Println("The length of array1 is:", len(array1))
        10 fmt.Println()
        11
        12 // adding values into the array
        13 fmt.Println("Adding values to array1")
        14 array1[0] = 29
        15 array1[1] = 17
        16 array1[2] = 42
        17 array1[3] = 13
        18 array1[4] = 56
        19 fmt.Println("Now the values of array1 are:", array1)
        20 fmt.Println()
        21
        22 // Reading all values from array1 using for-range loop
        23 fmt.Println("Reading all the values from array1 using for-range loop...")
        24 for index, value := range array1 {
        25 fmt.Printf("Index: %d Value: %d\n", index, value)
        26 }
        27 fmt.Println()
        28
    ```

_Review_

On line 7 we declare an array of type int with a length of 5 and print out the empty array.

    var array1 [5]int

On line 9 we print out the length of the array using a built-in function "len".

    fmt.Println("The length of array1 is:", len(array1))

From line 14 to 18 we initialize values to the array and then finally print out the filled array. The final array will look like the below image.

![Array Index](../images/golang-array-octallium.jpg)

> Did you notice that the index position starts from "0" and not from "1" !

So an array of length 5 will have index positions of 0, 1, 2, 3 & 4.

On line 24 we are using a "for-range" loop to print out the values, notice that we declare two variables "index" and "position" and then range over the array.

    for index, value := range array1

What this does is, it iterates over each element one by one and prints out the value & index, we don't have to declare variables and iterate, remember we used to write:

for i:= 0; i < n; i++ {
do something
}

Instead we can write it with much cleaner syntax using for-range loop, for-range loops are extensively used in Go. If you don't want to use any value, you can simply discard it using "\_" , for example if you do not want index position, you can declare for-range as:

    for _, value : range array1 {
        do something
    }

The underscore "\_" simply tells the Go compiler to ignore the value.

2.)

???+example "Part-2 array.go"

    ``` go
        29    // declaring an array of type string with shorthand method
        30    cities := [5]string{"New York", "Seattle", "Mumbai", "Sydney", "Montreal"}
        31    fmt.Println("The values of cities array are:", cities)
        32    fmt.Println()
        33
        34    // Reading specific values
        35    fmt.Println("Reading values from cities array using index position...")
        36    fmt.Println("The first city is:", cities[0])
        37    fmt.Println("The 3rd city is:", cities[2])
        38    fmt.Println("The last city is:", cities[4])
        39    fmt.Println("The last city can also be found by:", cities[len(cities)-1])
        40    fmt.Println()
        41
        42    // changing values
        43    fmt.Println("Changing the value at cities[0] to Paris...")
        44    cities[0] = "Paris"
        45    fmt.Println("The value at cities[0] is:", cities[0])
        46 }
    ```

_Review_

On line 30 we declare and initialze the array at the same time

    cities := [5]string{"New York", "Seattle", "Mumbai", "Sydney", "Montreal"}

We can access values at select index, checkout the lines from 36 to 39.

    fmt.Println("The first city is:", cities[0])
    fmt.Println("The 3rd city is:", cities[2])
    fmt.Println("The last city is:", cities[4])
    fmt.Println("The last city can also be found by:", cities[len(cities)-1])

We can also change the values in an array, on line 44 we change the value of cities[0] from New York to Paris.

    cities[0] = "Paris"

## Full Code

??? example "array.go"

    ``` go
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
    ```

## Running your code

Open your terminal and navigate to our folder

    code/basic/11_array/

Once in the folder type the following command

    go run array.go

## Output

If there are no errors, you should get the output as:

???+ success "Output"

        The values of array1 are: [0 0 0 0 0]
        The length of array1 is: 5

        Adding values to array1
        Now the values of array1 are: [29 17 42 13 56]

        Reading all the values from array1 using for-range loop...
        Index: 0 Value: 29
        Index: 1 Value: 17
        Index: 2 Value: 42
        Index: 3 Value: 13
        Index: 4 Value: 56

        The values of cities array are: [New York Seattle Mumbai Sydney Montreal]

        Reading values from cities array using index position...
        The first city is: New York
        The 3rd city is: Mumbai
        The last city is: Montreal
        The last city can also be found by: Montreal

        Changing the value at cities[0] to Paris...
        The value at cities[0] is: Paris

If for some reason your code isn't working, checkout the github repo.

## Github

[Github Repo](https://github.com/octallium/golang-handbook/tree/master/code)

## Golang Playground

[Golang Playground](https://play.golang.org/p/wZBg9UDUeDv)

## Limitation

Imagine you data is increasing and now you have 100 cities instead of 5, but the array can only hold 5 values, to overcome this limitation Go ships with a flexible type called as "Slice".

## Next

We will study {++ slices++} in depth in the next chapter.

## Please Donate ❤️

All the work is provided free of cost and completely open source, but it needs your support and love to keep the activity sustainable.

Any support is genuinely appreciated, you can help by sending a small donation by clicking the below link:

[![PayPal](../images/paypal-logo.png)](https://www.paypal.me/octallium)
