# {++ Shorthand++} Declaration Of Variables.

<hr>

## Objective

Learn how to declare variables using the shorthand method.

## Declaration

Go ships with a smart compiler, it can detect the data type and automatically assign it to variables, using short declaration is very widely used in Go, with this method you can create variables on the fly, no need for prior declaration.

General syntax

    variableName := value

It is also called as inference type declaration, meaning that the variable type is "inferred" from the value. This method offers various benefits such as compiler can choose the right data type and much more which we will explore as we write more code.

## Structure

Navigate to our code folder

    code/basic/

For our program create a new folder '06_shorthand_declaration'

    code/basic/06_shorthand_declaration

And lets create a file 'shorthand_declaration.go' in it, finally the structure would look like this:

    code/basic/06_shorthand_declaration/shorthand_declaration.go

## Code

???+ example "Shorthand Declaration"

    ``` go
    1 package main
    2
    3 import "fmt"

    4 func main() {
    5   // declaring integer
    6   num := 12
    7   fmt.Printf("The type of variable 'num' is: %T.\n", num)
    8
    9   // declaring float
    10  decimal := 15.45
    11  fmt.Printf("The type of variable 'decimal' is: %T.\n", decimal)
    12
    13  // declaring string
    14  name := "Octallium"
    15  fmt.Printf("The type of variable 'name' is : %T.\n", name)
    16 }
    ```

## Code Review

On line 6, 10 & 14 we declare a new variable

    num := 12
    decimal := 15.45
    name := "Octallium"

To check the data type we use a special format output function

    fmt.Printf()

Note, in the earlier examples we had used

    fmt.Println()

On line 7, 11 & 15, we print out the data type, to check the data type we use a special character "%T", which acts as a placeholder and represent the data "Type", it is followed by the variable name.

    fmt.Printf("The type of variable 'num' is: %T.\n", num)
    fmt.Printf("The type of variable 'decimal' is: %T.\n", decimal)
    fmt.Printf("The type of variable 'name' is : %T.\n", name)

If you don't understand the print statements, no worries, we will be having a dedicated section on formatting output, for now type everything as in the code above and make sure it runs.

## Run Code

Open your terminal and navigate to our folder

    code/basic/06_shorthand_declaration

Once in the folder type the following command

    go run shorthand_declaration.go

## Output

    The type of variable 'num' is: int.
    The type of variable 'decimal' is: float64.
    The type of variable 'name' is : string.

## Github

Just in case you have some errors with your code, you can check out the code at github repo

[Github Repo](https://github.com/octallium/golang-handbook/tree/master/code)

## Golang Playground

You can also run the code at playground

[Golang Playground](https://play.golang.org/p/9RLW4JSSwhq)

## Next

In the next chapter we will see about {++ if/else++} condition, don't worry if its becoming too geeky, keep up with the code and in no time you will get the hang of it.

## Please Donate ❤️

All the work is provided free of cost and completely open source, but it needs your support and love to keep the activity sustainable.

Any support is genuinely appreciated, you can help by sending a small donation by clicking the below link:

[![PayPal](../images/paypal-logo.png)](https://www.paypal.me/octallium)
