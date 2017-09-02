package main

import "fmt"

func main(){
    var x int = 32;
    fmt.Printf("The value of x is %d\n",x)
    fmt.Print("The address of x is ",&x)
    fmt.Print("\n")
    fmt.Printf("The Original memory of x is %d\n",&x)
}

/* Notes:

1) We initialize the value of x to 32.

2) We determine the memory address of x by giving "&x".

3) &x gives the memory address in Hexadecimal format.

4) We can get the Original memory of x with the help of %d in the Print statement.

*/
