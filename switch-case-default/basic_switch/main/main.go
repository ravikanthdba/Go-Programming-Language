package main

import "fmt"

func main() {
    switch "Ravikanth" {
        case "XYZ" : fmt.Println("XYZ")
        case "ABC" : fmt.Println("ABC")
        case "Ravikanth" : fmt.Println("Ravikanth")
        case "LMN" : fmt.Println("LMN")
        default : fmt.Println("Option doesn't exist")
    }
}


/* Notes:

1. The above program shows using of a switch statement.

2. The syntax for a switch statement:

    switch condition {
        case value1 : operation
        case value2 : operation
        case value3 : operation
        case value4 : operation
        case value5 : operation
        default     : operation
    }

3. We do not need to specify a break statement explicitly, as per other programs. The break is added by default.

4. If you need the next statement also to be printed, then we add a statement called "fallthrough"

5. If the condition doesn't match to any of the case statements, default is executed.

6. The "default" is optional.

*/
