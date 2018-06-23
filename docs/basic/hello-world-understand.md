# Understanding {++ "Hello, World!"++} Program.

<hr>

## Objective

Understanding the Hello World program line-by-line

## Code

???+ example "Hello World"

    ``` go
    1 package main
    2
    3 import "fmt"
    4
    5 func main() {
    6   fmt.Println("Hello, World!")
    7 }
    ```

## Code Review

Now its time to dive in and understand the code

Line 1 - 

    package main

Declares that this file belongs to the main package, in golang all the files are arranged in packages, we have a entire chapter dedicated to {++ Packages++} in future, for now remmember all files at your project directory level belong to package main.

Line 3 -

    import "fmt"

Here we say, they golang I want to import a package called as "fmt", "fmt" stands for "format".

"fmt" belongs to the golang standard library, means that the Go authors have already written it for you. You can simply import then and use all the code written by experts.

Documentation for "fmt" can be found [here](https://golang.org/pkg/fmt/)

"fmt" exposes a lot of functionality to us which we use in the line 6.

Line 5 - 

    func main () {

This is the start or entry point for your code, the compiler will automatically check if {++ func main()++} is present or not and start the program from there, it is generally mentioned as

    main.main
    [package].[function]

means "main" function in "main" package

Line 6 - 

    fmt.Println("Hello, World!)

Here we use Println function defined in the package "fmt" for printing out "Hello, World!" on the screen.

Now, you can see how we can use functionality written in packages and use it for our ease.

Line 7 - 

    }

Closes the main function, this symbolizes that main function ends here.

## Aside

Wow.. you have now learned your first Go program, even if all of the above doesn't make sense to you right now, don't worry ! You will soon be comfortable with all of that as we progress and write more code.

As a beginner, you would want to quit when you don't understand anything, my advise, don't quit just move along and later come back to topics which you didn't understand.

Learning code takes patience and persistance, fasten your seat belts and keep learning.
