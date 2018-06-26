# How To Add A {++ Comment++}.

<hr>

## Objective

Learn how to add a comment in the code. 

## Requirement

Adding comments is a very handy way of adding documentation to your code, it serves as a tool to let others know what you have written and makes navigating your code an easy task.

Go has a specific format for adding comments and recommends adding comments wherever it makes sense.

## Code

Adding a comment is pretty easy

    // This is a comment

The compiler ignores the comment and runs your code, however as humans for us to remmember what we have written and for others to understand writing comments would make sense.

## Hello World With Comment

???+ example "Hello World"

    ``` go
    1 // Declaring that the file belongs to the main package
    2 package main
    3
    4 // Importing the package "fmt"
    5 import "fmt"
    6
    7 // Declaring the main function
    8 func main() {
    9   
    10 // Printing Hello World to the standard output
    11 fmt.Println("Hello, World!")
    12 }
    13 // main function ends
    ```
## Golang Playground

???+ info "Hello World With Comment"

    Click on the below link:

    [Hello World With Comments](https://play.golang.org/p/Gy8TKpeAIO2)

## Next

Now that you know how to write comments, lets start with writing code with beautiful comments, but before we actually start writing more code there is one small step, we need to understand {++ Data Types++}

Click on the next section to understand basic data types in Go.