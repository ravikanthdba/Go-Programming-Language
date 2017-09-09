package main

import "fmt"

var integer int = 120

func main() {
    switch integer > 100 {
        case true : fmt.Printf("Yes, %d is greater than 100\n",integer); fallthrough;
        case false : fmt.Printf("Yes, %d is less than 100\n",integer)
        default : fmt.Println("No option")
    }
}


/* Notes:

1. The above program denotes usage of "Fallthrough"

2. In the above program, there is an integer, whose value is assigned to 120.

3. The switch checks for the condition if the integer is greater than 100.

4. The case has 3 conditions : true, false or default.

5. If its true, then it would print integer greater than 100.

6. Now, since "fallthrough" is also specified, the print in the next condition is also executed.

7. Hence it will also print "integer is less than 100".

8. With fallthrough, the next statement in the next condition, even if the condition i false will be executed.

9. If the value of integer is less than 100, then "integer is less than 100" will be printed.

10. Usually in a boolean login, only true or false is specified, but if there is another value, then the print statement in the default condition will be printed, but this is a rare scenario with respect to the above program.

*/
