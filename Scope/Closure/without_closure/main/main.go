package main

import "fmt"

var x int = 42;

func increment() int {
    fmt.Printf("The value of x while entering the increment function is : %d\n",x)
    x += 1
    return x
}


func main() {
    fmt.Printf("The initial value of x is : %d\n",x)
    fmt.Println(increment());
    fmt.Printf("The value of x after the first increment is %d\n",x)
    fmt.Println(increment());
}


/* Notes : The initial value of x is 42, as per the public scope.

In the main function, initially the value of x is set to 42.
We execute the increment function, and the value of x becomes 43.
The next time when we trigger the increment, the value of x which will be passed is 43. The value of x will not again reset to 42.
The second increment of x would give us 44

Therefore the output becomes :

1st Increment : 43
2nd Increment : 44

*/
