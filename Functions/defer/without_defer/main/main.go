package main

import "fmt"

func hello() {
    fmt.Println("hello")
}


func world() {
    fmt.Println("world")
}


func main() {
    world()
    hello()
}

/* Notes

1. In the above mentioned program, there are two functions:

	hello() and world()

hello() - prints the string "hello"
world() - prints the string "world"

2. In the main function, we have world() and hello() called in the above order.

3. Hence we have world and hello printed in sequential order:

world
hello

*/
