package main

import "fmt"

func modify(z *int){
    fmt.Printf("While entering: z = %d\n", *z)
    fmt.Printf("While entering: Address of z = %d\n", z)
    *z = 0
    fmt.Printf("While exiting: z = %d\n", *z)
    fmt.Printf("While exiting: Address of z = %d\n", z)
}

func main(){
    x := 42
    fmt.Printf("Before modifying: x = %d\n",x)   
    fmt.Printf("Before modifying: Address of x = %d\n",&x)   
    modify(&x)
    fmt.Printf("After modifying: x = %d\n",x)   
    fmt.Printf("After modifying: Address of x = %d\n",&x)   

}

/* Notes:

1. In the above program, we assign a value of x to 42. We are trying to modify the value of x through a function, and print the value of x after modification.

2. After assignment, we Print the value of x , which will return 42, and also print the address which is for example : 1234567890

3. We then call the modify function, which passes the memory address of x as an argument to the function.

4. Function modify takes the pointer as an input.

5. While entering the function, the value of z, which is *z, will be 42.

6. The memory address of z will be 1234567890, which is the same as given for x in the main function.

7. We change the value of *z, which implies *z => x (Please refer to basic Pointers if you do not understand it.) to 0

8. After changing, the value of *z becomes 0.

9. The memory address of z becomes : 1234567890.

10. In the main function, after modification, the value of x becomes 0.

11. The memory address of x will be 1234567890.

12. Since we are passing the address itself, instead of passing the value of x, the function was able to recognise, which value to be changed, hence after modification, we see the modified value of x, which satisfies our requirement.

*/
