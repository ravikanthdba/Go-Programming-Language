package main

import "fmt"


const name string  = "Ravikanth Garimella"
const x = 64

func main(){
    const y = 42;
    fmt.Printf("My name is %s\n", name)
    fmt.Printf("The value of y is %d\n", y)
    fmt.Printf("The value of x is %d\n", x)
}


/* Notes:

A constant means that never changing value. In the example, the value of "name" is never unchanging, and hence defined it as a constant and assigned the value "Ravikanth Garimella" to it. */
