package main

import "fmt"

var total,count int;


func calculate_average(numbers []int, callback func (n int) int) int {
    for _,n := range numbers {
        total += n
        count += 1
    }
    return total/count
}


func main() {
    average := calculate_average([]int{10,20,30,40,50,60},func (n int) int{
        return n
    })
    fmt.Println(average)
}
