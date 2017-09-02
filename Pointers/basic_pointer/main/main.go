package main

import "fmt"

func main() {
    var x int = 42;
    var str string = "Hello World"
    fmt.Printf("The value of x is %d\n",x)
    fmt.Print("The address of x is ",&x)
    fmt.Print("\n")
    fmt.Printf("The actual address of x is %d \n",&x)
    fmt.Print("\n\n\n")
    fmt.Printf("The value of str is %s\n",str)
    fmt.Print("The address of str is ",&str)
    fmt.Print("\n")
    fmt.Printf("The actual address of str is %d \n",&str)
}

/* Notes:

1. Pointer points to the memory address of the variable.

2. The memory address is given by "&".

3. The address would be an hexa-decimal value.

4. You can convert into integer using %d while printing the value of the memory.

5. The memory address varies for each and every run.

*/
