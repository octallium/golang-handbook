# Conditions {++ if/else++}.

<hr>

## Objective

To understand if/else conditions. 

If/else conditions are of very fundamental importance to programming and are present in all the programming languages.

The intent of condition is very simple

    if condition {
        do something
    } else {
        do another thing
    }

## Structure

Navigate to our code folder

    code/basic/

For our program create a new folder '07_if_else'

    code/basic/07_if_else

And lets create a file 'if_else.go' in it, finally the structure would look like this:

    code/basic/07_if_else/if_else.go

## Code

The code will be divided into two parts

1.)

???+example "if/else condition"
    ``` go
    1 package main
    2 
    3 import "fmt"
    4
    5 func main() {
	6 d := "Dog"
	7 c := "Cat"
    8
	9 // checking the value of variables
	10 if d == "Dog" {
	11	fmt.Println("Woff")
	12 } else {
	13	fmt.Println("I don't know which animal!")
	14 }
    15
    ```

Review

on line 10 we check if the value of the variable "d" is equal to "Dog"

    if d == "Dog"

If the condition is true then we print out "Woff" 

    fmt.Println("Woff")

If the condition is false, we print "I don't know which animal"

    fmt.Println("I don't know which animal!")


2.) If/else statements can also be chained if you have multiple conditions

???+example "if/else condition"

    ``` go
    16 // You can also chain if / else conditions
	17 if c == "monkey" {
	18	fmt.Println("I am a monkey.")
	19 } else if c == "Dog" {
	20	fmt.Println("I am a dog.")
	21 } else if c == "Cat" {
	22	fmt.Println("Meoww")
	23 }
    24 }
    ```

Review

On line 17 we check if value of the variable "c" is "monkey", if the conditions evaluates to true then we print "I am a monkey"

    if c == "monkey"

If it evaluates to false then we check it once again if it contains the value of "Dog"

    if c == "Dog"

Since, this also evaluates to false, we check for the next condition

    if c == "Cat"

As it evaluates to true, we print out "Meoww" on the screen

    fmt.Println("Meoww")

In case if "c" does not evaluate to true in any of the case, {++ nothing++} will be printed.

## Output

    Woff
    Meoww

## Note

Strings in Go are case sensitive, "monkey" and "Monkey" are evaluated differently, so be sure of using the right case when checking for evaluation.

## Github

Just in case you have some errors with your code, you can check out the code at github repo

[Github Repo](https://github.com/octallium/golang-handbook/tree/master/code)

## Golang Playground

You can also run the code at playground

[Golang Playground](https://play.golang.org/p/jd9it7OexoO)

## Next

We will see {++ for++} loops.