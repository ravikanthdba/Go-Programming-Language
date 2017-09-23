package main

import "fmt"


func example_callback(numbers []int,callback func(int)) {
    for n := range numbers {
        callback(n)
    }
}


func main() {
    example_callback([]int{1,2,3,4,5},func (n int) {
        fmt.Println(n)
    })
}
