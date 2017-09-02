package main

import "fmt"

func main(){
    var x int = 44;
    fmt.Printf("The value of x is %d\n", x)
    fmt.Printf("The address of x is %d\n", &x)
    fmt.Print("The address of x is ", &x)
    fmt.Printf("\n")
    var y *int = &x
    fmt.Printf("The value of y is %d\n",y)
    fmt.Print("The address of y is ",&y)
    fmt.Printf("\n")
    fmt.Printf("The value of *y is %d\n",*y)
    // fmt.Print("The value of **y is ",**y)
    fmt.Printf("\n")
    fmt.Printf("Change the value through the address:\n")
    *y = 84
    fmt.Printf("The value of x is :%d\n", x)
    fmt.Printf("The latest address of x is %d\n",&x)
}


/* Notes:

1. Variable x is initialized to 44.

2. The first Print statement prints the value of x, which is 44.

3. The second Print statement prints the address of x in integer format.

4. The third Print statement prints the address of x in Hexadecimal format.

5. The variable y is declared as a Pointer, and is assigned with address of x.

6. The 4th Print statement prints the value of y, which will be address of x.

7.  The 5th Print statement prints the address of y.

8. The 6th Print statement, prints *y, which returns the value of x.

9. Now through *y, we are changing the value of x to 84. As we see from above statement, *y points to the value of x, which means, if we change *y, it implies, we are changing x.

10. After changes, we try printing x, it gives us the modified value of x.

11. Even after changes, the address of x remains same, which means, the value under that address is overwritten with the new value.

12. We can pass memory address instead of passing the bunch of values, but still make the changes. This makes our programs very performant.

13. We do not need to pass large amounts of data, instead we can pass addresses.

14. Everything is pass by value in Go. Even if we pass the address, it will be pass by value itself and not pass by reference.

*/
