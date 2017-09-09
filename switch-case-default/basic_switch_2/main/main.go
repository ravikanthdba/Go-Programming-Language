package main

import "fmt"

func main() {
    integer := 190
    switch true {
        case integer > 180  : fmt.Printf("%d > 180\n",integer)
        case integer < 180  : fmt.Printf("%d < 180\n",integer)
        case integer == 180 : fmt.Printf("%d = 180\n",integer)
        case integer != 180 : fmt.Printf("%d != 180\n",integer)
        default : fmt.Printf("Invalid number\n")
    }
}
