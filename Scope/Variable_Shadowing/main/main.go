package main

import "fmt"

func max(number int) int {
    return 42 + number
}


func main() {
    max := max(7)
    fmt.Printf("The value of max is %d\n",max)
    // max := max(8)
    // fmt.Printf("The value of max is %d\n",max)
}


/* This is an example of bad programming. In the above program, we are declaring a global function called max, which can be used by any other function. In the main function, we are passing an argument 7 to the max function. In the max function, the variable will be assigned with 7, which will be only internal to that function. In the max function had computed 42 + 7, and returned the result. But in the main function, the returned result is stored again in the variable max. Now if we want to re-use the max function again, it would not pick up the function, its instead picking up the local variable which is defined. This is called variable shadowing, and this should never be used. */
