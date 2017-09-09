package main

import "fmt"

func main() {
    if 1 == 1 {
        if 1 == 2 {
            fmt.Printf("Nested if condition\n")
        } else {
            fmt.Printf("Coming out of Nested if condition\n")
        }
    } else {
            fmt.Printf("Coming out of the condition\n")
    }

}
