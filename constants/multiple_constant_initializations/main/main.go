// This program shows how to declare multiple constants.

package main

import "fmt"

// The following is the first way to define multiple constants and the best way.

const (
    name = "Ravikanth Garimella"
    val_x = 40
)


// The following is the second way to define constants

const myname string = "Hello World"
const val_y int = 50

func main() {
    fmt.Printf("My name value through 1st method is : %s\n", name)
    fmt.Printf("The value of x through 1st method is : %d\n", val_x)
    fmt.Printf("My name value through 2nd method is : %s\n", myname)
    fmt.Printf("The value of y through 2nd method is : %d\n", val_y)
}
