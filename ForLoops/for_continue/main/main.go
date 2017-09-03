package main

import "fmt"

func main() {
    i := 0
    for {
        i ++
        if i % 2 == 0 {
            continue
        } else {
             fmt.Println(i)
        }
        if i > 100 {
            break 
        } 
    }
}

/* Notes:

1. This program shows how to use "Continue"

2. We initialize the value of i to 0.

3. We start the for loop.

4. First increment the value of i. So i becomes 1.

5. then we check 1 % 2 = 1. Since it is not equal to 0, it prints 1.

6. It then checks if 1 > 100, which is false, so it goes to the starting of the loop again.

7. Increment the value of i to 2.

8. 2 % 2 = 0, so it doesn't do anything. the continue statement moves the cursor to starting of the for loop.

9. i will be incremented, and new value of i would be 3.

10. This would continue until it reaches 100.

11. Upon reaching 100, since 100 % 2 = 0, it continues, cursor moves to starting of for loop. i will be incremented to 101.

12. since 101 % 2 = 1, it Prints 101, and checks the condition. since 101 > 100, it comes out of the loop.

13. Therefor it prints all odd numbers between 0 and 102.

*/


