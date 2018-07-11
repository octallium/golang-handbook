# Understanding {++ For Loops++}.

<hr>

## Objective

To understand {++ for++} loops.

Go offers only "for" loops for doing repetitive tasks, if you come from some other languages there are no "while" or "do-while" loops in Go, instead "for" loops have a flexible syntax and offers all the looping functions.

General syntax is:

    for condition {
        do something
    }

## Structure

Navigate to our code folder

    code/basic/

For our program create a new folder '08_for_loops'

    code/basic/08_for_loops

And lets create a file 'for_loops.go' in it, finally the structure would look like this:

    code/basic/08_for_loops/for_loops.go

## Code

The code will be divided into four parts:

1.)

???+example "for loops"
    ``` go
    1 package main
    2
    3 import "fmt"
    4
    5 func main() {
	6 // Variation 1 - conditional loop
	7 num1 := 0
	8 fmt.Println("Starting for loop...")
	9 for i := 0; i < 11; i++ {
	10  fmt.Println("Num =", num1)
	11	num1++
	12 }
	13 fmt.Println()
    ```

_Review_

Line 9 declared a for loop along with condition

    for i := 0; i < 11; i++ {

We initial the variable "i" and set it to "0", then we say loop till "i is less than 11" and after each pass do "i++", i.e increment the value of "i" by 1.

Line 10 prints the value after each pass and on line 11 we increment the value of num

    fmt.Println("Num =", num1)
    num1++

2.)

???+example "for loops"
    ``` go
 	14 // Variation 2 - Infinite loop
    15 num2 := 20
	16 fmt.Println("Entering infinite loop...")
	17 for {
	18	// break condition
	19	if num2 < 10 {
	20		break
	21	}
	22	fmt.Println("Num =", num2)
	23	num2--
	24 }
	25 fmt.Println()
    ```

_Review_

On line 17 we start a infinite loop, the for loop will keep on executing till it encounters a break condition or runs out of memory

    for {

On line 19 we check for a break condition

    if num2 < 10

3.)

???+example "for loops"
    ``` go
 	26	// Variation 3 - Optional statements
	27 num3 := 20
	28 fmt.Println("Loop with optional statements...")
	29 for num3 <= 30 {
	30	fmt.Println("Num =", num3)
	31	num3++
	32 }
	33 fmt.Println()
    ```

_Review_

On line 29 we start a loop with operational statements, the loop will break when it fulfills the condition.

    for num3 <= 30

Note we have declared and initialized "num3" on line 27

    num3 := 20

4.)

???+example "for loops"
    ``` go
 	34	// Variation 4 - Boolean operators
	35 cond := true
	36 num4 := 30
	37 fmt.Println("Loops with boolean operator...")
	38 for cond {
	39	if num4 >= 40 {
	40		cond = false
	41	}
	42	fmt.Println("Num =", num4)
	43	num4++
	44 }
    45 }
    ```

_Review_

We declare a boolean condition on line 35, and start the loop on line 38

    for cond

In order to break the loop we set the condition on line 39

    if num4 >= 40 {
		cond = false
	}

## Full Code

??? example "Complete For-Loop Code"
    ``` go
    package main

    import "fmt"

    func main() {
	    // Variation 1 - conditional loop
	    num1 := 0
	    fmt.Println("Starting for loop...")
	    for i := 0; i < 11; i++ {
		    fmt.Println("Num =", num1)
		    num1++
	    }
	    fmt.Println()
	    // Variation 2 - Infinite loop
	    num2 := 20
	    fmt.Println("Entering infinite loop...")
	    for {
		    // break condition
		    if num2 < 10 {
			    break
		    }
		    fmt.Println("Num =", num2)
		    num2--
	    }
	    fmt.Println()
	    // Variation 3 - Optional statements
	    num3 := 20
	    fmt.Println("Loop with optional statements...")
	    for num3 <= 30 {
		    fmt.Println("Num =", num3)
		    num3++
	    }
	    fmt.Println()
	    // Variation 4 - Boolean operators
	    cond := true
	    num4 := 30
	    fmt.Println("Loops with boolean operator...")
	    for cond {
		    if num4 >= 40 {
			    cond = false
		    }
		    fmt.Println("Num =", num4)
		    num4++
	    }
    }
    ```

## Note

Go ships with one more variant know as "for - range", we will study it in the coming chapters.

## Github

Just in case you have some errors with your code, you can check out the code at github repo

[Github Repo](https://github.com/octallium/golang-handbook/tree/master/code)

## Golang Playground

You can also run the code at playground

[Golang Playground](https://play.golang.org/p/LC7FfsSdaHB)

## Next

In the next chapter we will learn about {++ function++} declaration.