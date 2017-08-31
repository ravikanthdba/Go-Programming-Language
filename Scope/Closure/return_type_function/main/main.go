package main

import "fmt"

func assign_value() func() int {
    x := 42;
    return func() int{
        x++;
        return x;
    }
}

func main() {
    increment := assign_value();
    fmt.Println(increment);
    fmt.Println(increment());
    fmt.Println(increment());
}


/* Notes:

The above program shows, we can even return the Functions from the functions.

We have implemented "Closure" to the above program

1) We have created a fucntion assign_value, which is returning a function in return.

2) In the function assign_value, we have assigned a value x to 42.

3) In the function main,increment has been assigned a value, by calling the function assign_value().

4) When we print the value increment, it will be the pointer to the function.

5) We are incrementing the value twice

*/

