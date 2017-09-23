package main

import "fmt"

func main() {
    data := []int{3,6,7,4,10}
    avg_cal := calculate_avg(data...)
    fmt.Printf("The average of numbers is %d\n",avg_cal)
}

func calculate_avg(input ...int) int {
    fmt.Println("The input dataset is ",input)
    var total int
    for _,val := range input {
        total += val
    }
    return total / len(input)
}
