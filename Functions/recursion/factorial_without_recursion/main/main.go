package main

import "fmt"

func main() {
    n := 4
    var factorial int = 1
    for n != 0{
        factorial *= n
        n -= 1
    }
    fmt.Printf("The factorial is %d\n",factorial)
}
