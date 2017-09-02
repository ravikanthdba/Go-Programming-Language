package main

import "fmt"

func main() {
    x := 13 % 2
    fmt.Printf("The remainder of 13 and 2 is %d\n",x)

    for i := 1; i <= 100; i++ {
        if i % 2 != 0 {
            fmt.Printf("%d is Odd\n",i)
        } else {
            fmt.Printf("%d is even\n",i)
        }
    }

}

/* Notes:

1. This Program is mainly used for finding the remainder after dividing 2 variables.

2. Remainder can be found with the operation : %

3. Initially, we are trying to find the remainder of 13 and 2, which is 1, and printing the remainder value.

4. We are then looping from 1 to 100 variables, and figuring out which one is odd and which one is even.

5. Very important syntactical statement : when we write an if, as and when the if block closes, the else has to be immediately next to the closing parenthesis of the if block. It should not be in the new line. The reason is by default go assigns a ";" semi-colon as and when a statement is complete. Hence, if we write else in a new line, it is being treated differently, and hence shows an error message.

*/
