package main

import "fmt"

func filtered_records(numbers []int,callback func(int) bool) []int {
    xs := []int{}
    for n := range numbers {
        if callback(n) { xs = append(xs,n)}
    }
    return xs
}


func main() {
    filter := filtered_records([]int{1,2,3,4,5},func (n int) bool{
        return n > 1
    })
    fmt.Println(filter)
}
