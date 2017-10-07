package main

import "fmt"

func factorial(n int) int {
    if n == 0 { 
        return 1
    }
    return n * factorial(n-1)
}


func main() {
    x := 3
    fmt.Printf("The factorial of %d is %d\n",x,factorial(x))
}
