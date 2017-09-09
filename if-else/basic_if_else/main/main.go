package main

import "fmt"

func main() {
    if true {
        fmt.Printf("This works\n")
    }    else {
        fmt.Printf("This does not print\n")
    }
}

/* NOTES:

1) The above program shows a basic if-else program.

2) The syntax of if-else is :

    if condition {
        operation
    } else {
        operation
    }

3) The else has to  be right after the "If" bracket has ended as shown above, else Go adds a ";" by default after every line, else, when it starts in a new line, will be considered separate.

4) In the above condition, the default is true, and hence "This works" gets printed.

*/
