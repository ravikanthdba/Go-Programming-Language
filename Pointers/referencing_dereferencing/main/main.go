package main

import "fmt"

func main(){
    x := 42;
    fmt.Printf("The value of x is %d\n",x)
    var y *int = &x;
    fmt.Printf("The value of y is %v\n",y)
    fmt.Printf("The memory address of y is %v\n",&y)
    fmt.Printf("The value of *y is %v\n",*y)
    fmt.Printf("The value of &*y is %v\n",&*y)
//    fmt.Printf("The value of **y is %v",**y)
//    fmt.Printf("The value of &&y is ",&&y)
//    fmt.Printf("The value of &(*y) is ",&(*y))
//    fmt.Printf("The value of *(&y) is ",*(&y))

}

/* Notes:

1. The above program depicts referencing and dereferencing of the Pointer.

2. We initialize an integer variable x to 42.

3. We initialize a pointer variable y to memory location of x.

4. The value of x becomes 42.

5. The value of y will become address of x.

6. The value of *y becomes the value of x, which is 42.

7. The value of &y becomes the memory address of y.

8. The value of &*y becomes the address of x, as eventually &*y => &x, which is address of x.

9. The other values like **y, &&y , &*y and *&y throws an error message.

*/
