# Writing a {++ "Hello, World!"++} program.

<hr>

## Objective

Write a program which prints "Hello, World!" on the screen.

## Structure

Let's create a folder called 'code' anywhere on your machine and we will put all our Go code in it.
Inside the 'code' folder lets create one more folder for basic tutorials

    code/basic/

For our first program create a new folder '01_hello_world'

    code/basic/01_hello_world

And lets create a file 'hello_world.go' in it, finally the structure would look like this:

    code/basic/01_hello_world/hello_world.go

## Code

Write the code as shown below, while you can simply copy and paste, its better if you write everything on your own.

Don't worry if you code won't work, only then copy and paste the code :smile:

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

## Running your code

Open your terminal and navigate to our folder

    code/01_hello_world/hello_world.go

Once in the folder type the following command

    go run hello_world.go

## Build

Instead of just running let's try to build our code by compiling it to binary.

Instead of "go run" type the following command

    go build hello_world.go

If you are running on a unix based system including mac OSx, you can run the binary by typing

    ./hello_world

## Output

If there are no errors, you should get the output as:

???+ success "Output"

        Hello, World!

If for some reason your code isn't working, checkout the golang playground or github links in the following section.

## Github

[Github Repo](https://github.com/octallium/golang-handbook/tree/master/code)

That's it, {++ Congratulations++} ! You just wrote your first Go program.

## Golang Playground

Golang has a online sandbox environment for running your Go programs, which can be accessed on [Golang Playground](https://play.golang.org/)

I will be posting all the playground links for all the code we write, this way you can run them online and compare with your code.

???+ info "Hello World On Playground"

    Click on the below link:

    [Hello World](https://play.golang.org/p/RBSYyNyyVba)

## Next

If you haven't understood anything what you wrote, no worries, you aren't expected to understand it just yet !

In the next section we will understand everything line-by-line.

## Please Donate ❤️

All the work is provided free of cost and completely open source, but it needs your support and love to keep the activity sustainable.

Any support is genuinely appreciated, you can help by sending a small donation by clicking the below link:

[![PayPal](../images/paypal-logo.png)](https://www.paypal.me/octallium)
