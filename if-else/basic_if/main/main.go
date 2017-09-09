package main

import "fmt"

func main() {
    if true {
        fmt.Printf("This works\n")
    }

    if false {
        fmt.Printf("This does not work\n")
    }
}

/* Notes:

1. The above program shows example "IF" condition.

2. In the above program, the default value is True, and hence "This works is being printed."

3. If : always has a boolean condition, and based on the boolean output, its corresponding operation is performed.

4. The boolean value has to be "true" or "false", since Go is case in-sensitive, "True" is not equal to "true".

5. The syntax for basic if is :

	if condition {
            operation
        }

*/
