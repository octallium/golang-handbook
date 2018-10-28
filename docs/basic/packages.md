# Understanding {++ Packages++}.

<hr>

## Objective

To understand {++ packages++} in Go.

You have already used packages, remember we imported "fmt" package and used Println() function in it.

Packages help us in organizing our code, we can group together related functions in a package. You can create multiple functions and files inside a package.

Note, that there can be only one {++ func main()++} throughout your project, as main() is the starting point in any standalone application.

## Structure

Navigate to our code folder

    code/basic/

For our program create a new folder '10_packages'

    code/basic/10_packages/

And lets create a file 'main.go' in it, finally the structure would look like this:

    code/basic/10_packages/main.go

## Declaring Packages

We will be creating a custom package called as "greeting" and have a couple of functions in it, but we first need to declare a package.

Navigate to our directory

    code/basic/10_packages/

For declaring a package create a folder named "greeting", now the structure would be

    code/basic/10_packages/greeting

Inside the "greeting" folder create a file "greet.go", we will write our functions inside "greet.go" and call it from main() as required.

    code/basic/10_packages/greeting/greet.go

Now our structure would look like

![Package Structure](../images/packages-tree-octaliium.png)

## Code

A) First let us write the code for our custom package "greeting"

Open {++ greet.go++} from:

    code/basic/10_packages/greeting/greet.go

We will write the code in 4 parts:

1.)

???+example "Part-1 greet.go"
`go 1 package greeting 2 3 // The above statement declares that this file belong to 4 // the package "greeting" 5 6 // import other packages 7 import "fmt" 8 9 // GoodMorning - Note that the name of our function starts with a capital alphabet, 10 // capital alphabet denotes that the function is exported or visible outside 11 // the package. 12 // Now you can understand why Println() has a capital letter in the beginning ! 13 func GoodMorning() { 14 fmt.Println("Good Morning, Octallium!") 15 fmt.Println() 16 } 17`

_Review_

On line 1 we declare that this file belongs to the package "greeting", note that the package name and the parent directory name has to be the same.

    package greeting

On line 13 we declare a function "GoodMorning" which prints out "Good Morning, Octallium!" on the screen

    func GoodMorning()

Note that the first alphabet is capital, it symbolizes that the function is exported or visible to other packages, we will call it from our main().

2.)

???+example "Part-2 greet.go"
`go 18 // goodAfterNoon starts with a lowercase alphabet, means that this function is not 19 // exported, and can only be accessed within the package. 20 // Calling goodAfterNoon from another package will result in an error 21 func goodAfterNoon() { 22 fmt.Println("Good Afternoon, Octallium!") 23 fmt.Println() 24 } 25 26 // CallGoodAfterNoon has access to goodAfterNoon as it belongs to the same file, 27 // and hence it can call it. 28 func CallGoodAfterNoon() { 29 goodAfterNoon() 30 } 31`

_Review_

On line 21 we declare a function "goodAfterNoon", but since the first alphabet is lowercase, the function is not exported, i.e it is not visible from other packages and can be accessed only within the file.

    func goodAfterNoon()

In order to access it, on line 28 we create another function "CallGoodAfterNoon", it can access goodAfterNoon as it belongs to the same file and we can call goodAfterNoon() by calling CallGoodAfterNoon().

    func CallGoodAfterNoon()

3.)

???+example "Part-3 greet.go"
`go 32 // GoodNight accepts a string parameter and prints the result to the screen 33 func GoodNight(name string) { 34 fmt.Println("Good Night, ", name) 35 fmt.Println() 36 } 37`

_Review_

On line 33 we declare a function "GoodNight" which accepts a string and prints out a message. Note that the function doesn't return anything.

    func GoodNight(name string)

4.)

???+example "Part-4 greet.go"
`go 38 // Factorial accepts an int parameter and also returns an int, 39 // it calculates the factorial of the input int and returns 40 // the factorial. 41 // At this point ignore how the function works, just concentrate 42 // on the declaration and syntax. 43 func Factorial(n int) int { 44 if n == 0 { 45 return 1 46 } 47 return n * Factorial(n-1) 48 }`

_Review_

On line 43 we declare a function "Factorial" which accepts an int and also returns an int.

    func Factorial(n int) int

At this point don't think how the function is working, let's concentrate on the declaration and syntax.

???example "Complete greet.go"
``` go
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
    ```

B) Let's write the main.go

Navigate and open main.go from:

    code/basic/10_packages/main.go

???+ example "main.go"

    ``` go
    1 package main
    2
    3 // The above declaration says that this file belongs to the "main" package
    4
    5 // importing our packages
    6 import (
    7    // package "fmt" comes in-built as a part of Go standard library
    8    "fmt"
    9    // importing our custom package
    10    "./greeting"
    11    // the "./" denotes that from the current directory go to the "greeting" directory
    12  )
    13
    14 func main() {
    15    // Calling the GoodMorning() from greeting package
    16    fmt.Println("Calling function GoodMorning() from package 'greeting'...")
    17    greeting.GoodMorning()
    18    // Calling the goodAfterNoon() from greeting package
    19    fmt.Println("Calling function goodAfterNoon() from package 'greeting'...")
    20    // as goodAfterNoon is not directly accessible, we call CallGoodAfterNoon()
    21    // which in turn calls goodAfterNoon()
    22    greeting.CallGoodAfterNoon()
    23    // Calling the GoodNight() from greeting package
    24    fmt.Println("Calling function GoodNight() from package 'greeting'...")
    25    greeting.GoodNight("Gopher!")
    26    // Calling the Factorial function with passing an int
    27    num := 10
    28    // As Factorial() returns an int, we save it in a new variable "result"
    29    result := greeting.Factorial(num)
    30    // Printing the result on screen, note we are using Printf()
    31    fmt.Printf("Factorial of %d is: %d.\n", num, result)
    32 }

    ```

On line 1 we declare that the file belongs to the package "main"

    package main

On line 10 we import our custom package "greeting"

    "./greeting"

The "./" denotes that from the current directory go to the "greeting" directory. On line 17 we call the GoodMorning function

    greeting.GoodMorning()

Note the syntax

    packageName.funcName

If we try to directly call goodAfterNoon() we will get an error

    Error: greeting.goodAfterNoon()

So, we call it by calling the CallGoodAfterNoon() function on line 22.

    greeting.CallGoodAfterNoon()

On line 25 we call the GoodNight() function and pass a string "Gopher" along with it.

    greeting.GoodNight("Gopher!")

On line 29 we call the Factorial() function and pass an int, since it also returns an int, we save it in a variable "result" and print the output on line 31

    result := greeting.Factorial(num)

## Running your code

Open your terminal and navigate to our folder

    code/basic/10_packages/

Once in the folder type the following command

    go run main.go

## Build

Till, now we have been running our code, let's try to build our code by compiling it to binary.

Instead of "go run" type the following command

    go build main.go

If you are running on a unix based system including mac OSx, you can run the binary by typing

    ./main

If there are no errors, you should get the output as:

???+ success "Output"

        Calling function GoodMorning() from package 'greeting'...
        Good Morning, Octallium!

        Calling function goodAfterNoon() from package 'greeting'...
        Good Afternoon, Octallium!

        Calling function GoodNight() from package 'greeting'...
        Good Night,  Gopher!

        Factorial of 10 is: 3628800.

If for some reason your code isn't working, checkout the github repo.

## Github

[Github Repo](https://github.com/octallium/golang-handbook/tree/master/code)

That's it, {++ Congratulations++} ! You just wrote your first custom package in Go.

## Golang Playground

Since golang playground does not allow to declare and use custom packages, it is not possible to post it there, please refer the github repo for any errors.

## Next

We will start the basic data structures of type {++ array++}.

## Please Donate ❤️

All the work is provided free of cost and completely open source, but it needs your support and love to keep the activity sustainable.

Any support is genuinely appreciated, you can help by sending a small donation by clicking the below link:

[![PayPal](../images/paypal-logo.png)](https://www.paypal.me/octallium)
