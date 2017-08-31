package main

import "fmt"

func main() {
    var x int = 42;
    {
       fmt.Printf("The value of x is %d\n", x);
       var y int = 45;
       fmt.Printf("This is from the inner scope : The value of y is %d\n", y)
    }
  // fmt.Printf("The value of variable y is %d", y)
}


/* Variable main is the outer scope. The value x is defined in the outer scope. The braces inside the main function defines the inner scope. Any variable defined on the outer scope can be accessed in the inner scope. Any variable defined in the inner scope, is limited to that inner scope itself. It cannot be used in the outer scope.

The above program is similar to the following program:

package main

import "fmt"

var x int = 42;

func main(){
    fmt.Printf("The value of x is %d", x);
    foo();   
}

func foo() {
    var y int = 45;
    fmt.Printf("The value of y is %d", y)
    fmt.Printf("From the Function foo, The value of x is %d", x)
}


In the above example, The variable x is a global variable, and it can be accessed in all the functions, Where as the variable y is limited only to the function "foo"


Very Important Notes:

Closure helps us to limit the scope of variables used by multiple functions.
Without closure, if a variable needs to be accessed by two functions, the variable needs to be a package scope.

*/
