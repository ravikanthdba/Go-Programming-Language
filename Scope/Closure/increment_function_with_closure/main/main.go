package main

import "fmt"

func main() {
    var  x int = 43
    increment := func() int {
        x++ ;
        return x;
    }
 
 fmt.Println(increment());
 fmt.Println(increment());

}


/* Notes:

We have implemented "Closure" to the increment function.

So we have reduced the scope of using the variable. Now the variable x is not a package related variable.

When implementing a inner function, we implement an "Anonymous function"

Anonymous Function : The function without a name is called an Anonymous function

Func expression : Assigning a function to a variable.

*/
