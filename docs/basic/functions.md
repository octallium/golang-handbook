# Understanding {++ Functions++}.

<hr>

## Objective

To understand {++ functions++} in Go.

Functions are first class citizens in Go, in fact you have already been using functions, we have been declaring "main" function and using it from our first program.

General syntax is:

    func funcName(optional parameters) return parameters {
        do something
    }

## Structure

Navigate to our code folder

    code/basic/

For our program create a new folder '09_functions'

    code/basic/09_functions/

And lets create a file 'functions.go' in it, finally the structure would look like this:

    code/basic/09_functions/functions.go

## Code

The code will be divided into four parts:

1.)

???+example "Functions"

    ```go
        1 package main
        2
        3 import "fmt"
        4
        5 // sayHello prints "Hello, Octy!"
        6 func sayHello() {
        7 fmt.Println("In func sayHello()...")
        8 fmt.Println("Hello, Octy!")
        9 fmt.Println()
        10 }
        11
    ```

_Review_

Line 6 defines a function "sayHello" which prints "Hello, Octy!" on screen

    func sayHello()

2.)

???+example "Functions"

    ```go
        12 // function with passing values
        13 func sayHelloTo(name string) {
        14 fmt.Println("In func sayHelloTo()...")
        15 fmt.Println("Hello,", name)
        16 fmt.Println()
        17 }
        18
    ```

_Review_

On line 13 we declare a function "sayHelloTo" with accepts a parameter.

    func sayHelloTo(name string)

sayHelloTo accepts a string "name" as a parameter and then print it on line 15.

3.)

???+example "Functions"

    ```go
        19 // function with return values
        20 func printNums(n int) error {
        21 fmt.Println("In func printNums()...")
        22 var err error
        23 for i := 0; i <= n; i++ {
        24 _, err = fmt.Println(i)
        25 }
        26 return err
        27 }
        28
    ```

_Review_

Line 20 we define a function "printNums" with accepts a parameter and returns a parameter of type "error".

    func printNums(n int) error

On line 22 we define a variable "err" of type "error", till now we have seen data types of integer, float, boolean, string, similarly Go offers a custom type "error".

    var err error

Line 24 is a little special, Println() returns two values, "number of bytes written" and an "error", as we don't want number of bytes written we use a underscore "\_", Go compiler ignores underscore and throws away its value.

    _, err = fmt.Println(i)

We catch the error value in the variable "err" and return it on line 26.

    return err

4.)

???+example "Functions"

    ```go
        29 // main function
        30 func main() {
        31 fmt.Println("In func main()")
        32 fmt.Println("Now calling func sayHello()...")
        33 fmt.Println()
        34 // calling function sayHello()
        35 sayHello()
        36 // calling a function with passing values
        37 name := "Gopher"
        38 sayHelloTo(name)
        39 // calling a function with return parameters
        40 err := printNums(10)
        41 if err == nil {
        42 fmt.Println("There are no errors!")
        43 } else {
        44 fmt.Println("Error with printing:", err)
        45 }
        46 }
    ```

_Review_

We call individual functions on lines 35, 38 & 40.

    sayHello()

    sayHelloTo(name)

We pass a value "name" along with the function.

    err := printNums(10)

We call printNums along with a value "10" and assign it to the variable "err", the return error gets stored in the variable "err"

## Full Code

??? example "Complete Functions Code"

    ```go
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
    ```

## Running your code

Open your terminal and navigate to our folder

    code/basic/09_functions/

Once in the folder type the following command

    go run main.go

## Output

???+ success "Output"

        In func main()
        Now calling func sayHello()...

        In func sayHello()...
        Hello, Octy!

        In func sayHelloTo()...
        Hello, Gopher

        In func printNums()...
        0
        1
        2
        3
        4
        5
        6
        7
        8
        9
        10
        There are no errors!

## Github

Just in case you have some errors with your code, you can check out the code at github repo

[Github Repo](https://github.com/octallium/golang-handbook/tree/master/code)

## Golang Playground

You can also run the code at playground

[Golang Playground](https://play.golang.org/p/K-1VZCt_M9k)

## Next

In the next chapter we will learn about {++ packages++} in Go.

## Please Donate ❤️

All the work is provided free of cost and completely open source, but it needs your support and love to keep the activity sustainable.

Any support is genuinely appreciated, you can help by sending a small donation by clicking the below link:

[![PayPal](../images/paypal-logo.png)](https://www.paypal.me/octallium)

```

```
