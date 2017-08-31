package main

import "fmt"

// Since variable "x" is declared outside of all the functions, this variable is accessed by all the functions.


var x int = 45

func main() {
    fmt.Printf("The variable from the main function is %d\n",x)
    foo() 
}


func foo() {
    fmt.Printf("the value of the variable from the Foo function is %d\n",x)
}
