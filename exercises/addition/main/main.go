package main

import "fmt"

func main() {
    var x int ;
    var y int ;
    fmt.Println("Enter the values of x and y:")
    fmt.Scan(&x)
    fmt.Scan(&y)
    input_x := x
    input_y := y
    iteration := 1
    carry := 0
    sum_total := 0

    for x % 10 != 0 && y % 10 != 0 {
        addition := ((x%10) + (y%10) + carry)
        if addition > 9 {
            carry = addition / 10
        } else {
            carry = 0 
        }
        addition = addition % 10
        sum_total = sum_total + (addition * iteration)
        iteration = iteration * 10
        x = x / 10
        y = y / 10
    }
    sum_total = sum_total + ((x + y + carry) * iteration)
    fmt.Printf("The sum of %d and %d is %d\n",input_x,input_y,sum_total)
}
