package main

import "fmt"

func modify(z int){
    fmt.Printf("While entering the modify function, the value of z is %d\n",z)
    fmt.Printf("While entering the address of x is %d\n",&z)
    z = 0;
    fmt.Printf("While exiting the modify function, the value of z is %d\n",z)
    fmt.Printf("While exiting the address of x is %d\n",&z)
}


func main(){
 x := 5;
 fmt.Printf("Before modifying the value of x is %d\n",x)
 fmt.Printf("Before modifying the address of x is %d\n",&x)
 modify(x)
 fmt.Printf("After modifying teh value of x is %d\n",x)
 fmt.Printf("After modifying the address of x is %d\n",&x)
}


/* Notes:

1. In the above program, we assign a value of x to 5. We are trying to modify the value of x through a function, and print the value of x after modification.

2. After assignment, we Print the value of x , which will return 5, and also print the address which is for example : 1234567890

3. We call the modify function, which the argument of x. i.e. We pass the argument 5 into the modify function.

4. In the function, we call the variable as z, for no confusion.

5. When entering the function the value of z will be 5.

6. The address of z will be different, the address will become 1234567891.

7. We set the value of z to 0, i.e. ideally the value of x should become 0. But since the memory address of both the variables are different, there will be no change to the value, even after modification.

8. After modification, the value of z will be 0.

9. The memory address of z after changes will be 1234567891.

10. In the main function, after the modify function is called, the value of x will still be 5.

11. The memory address of x will still be 1234567890.

12. Hence the program doesn't satisfy our requirement.

13. With the help of Pointers, we would be able to solve this problem

*/
