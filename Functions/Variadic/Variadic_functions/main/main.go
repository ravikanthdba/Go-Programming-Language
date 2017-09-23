package main

import "fmt"

func main() {
    avg_list := cal_average(34,45,56,67,78)
    fmt.Printf("The average of the numbers is %d\n",avg_list)
}

func cal_average(avg_object ...int) int {
    fmt.Println("The list is ",avg_object)
    fmt.Printf("The type is %T\n",avg_object)
    var total int
    for _,val := range avg_object {
        total += val
    }
    return int(total) / int(len(avg_object))
}

/* Notes:

1. This program calculates the average of 5 numbers using Variadic functions.

2. Variadic functions are represented by "..."

*/
