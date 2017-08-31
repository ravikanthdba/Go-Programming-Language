/* Block level scope is limiting the variable to a particular function */

package main

import "fmt"

var x int = 45

func main() {
    fmt.Printf("Value of x from main is %d\n", x)
    foo()
    bar()
    y := 20
    fmt.Printf("Value of y from main is %d\n", y)
    // fmt.Println("Value of z from main is %d\n", z)
    // z := 14
    /* Ordering matters, we cannot print a particular variable and define it later. In the above example, variable "z" was printed first and declared later, and hence the program failed. */
}


func foo() {
    fmt.Printf("Value of x from Foo function is %d\n", x)
    // fmt.Printf("Value of y from Foo function is %d\n", y)

    /* Since the variable "y" is defined within the Parenthesis of the main function, it is not accessible by another function, since its scope is limited only to main. */

    fmt.Printf("Value of z from Foo is %d\n",z)
}


func bar() {
    fmt.Printf("Value of z from Bar is %d\n",z)
}

var z int = 99


/* In the above case, we could see, we have defined Two functions Foo() and Bar(), which is Printing the value "z", and z is defined after it. However since "z" is a package level variable, it can be given anywhere in the package, and all functions declared before or after the variable declaration would be able to Print the value of "z", as the scope of package is more than the scope of the block.*/
