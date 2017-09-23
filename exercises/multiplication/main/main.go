package main

import "fmt"

func main() {
    var x int;
    var y int;
    fmt.Println("Multiplication of 2 numbers:")
    fmt.Printf("Enter the value of x:")
    fmt.Scan(&x)
    fmt.Printf("\n")
    fmt.Printf("Enter the value of y:")
    fmt.Scan(&y)
    iteration := 1
    product := 0
    var inner_product int;
    for (y % 10 != 0) {
        var y_rem int;
        y_rem = y % 10
        inner_product = (x * y_rem) * iteration
        y = y / 10
        iteration = iteration * 10
        product = product + inner_product
    }
    fmt.Printf("The product of %d and %d is %d\n",x,y,product)
}
