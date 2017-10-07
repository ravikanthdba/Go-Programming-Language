package main

import "fmt"

var total int;

func sum_numbers (numbers[]int,callback func (n int) int) int {
    for _,n := range numbers {
        total += n
    }            
    return total
}


func main() {
    sum := sum_numbers([]int{80,20,10,30,40,80},func (n int) int {
        return n
    })
    fmt.Println(sum)
}
