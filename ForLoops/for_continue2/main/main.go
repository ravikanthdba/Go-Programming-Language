package main

import "fmt"

func main() {
    i := 0
    for {
        i ++ 
        if i > 100 {
            break
        }
        if i % 2 == 0 {
            continue
        } else {
            fmt.Println(i)
        }

    }
}

/* Notes:

1. The above program shows how "continue" works.

2. "break" can be accessible only in loop "for".

3. If you want to skip doing any operation on a condition and keep moving to the next variable, then we use a "continue" operation.

4. The above program we print odd numbers between 1 and 100.

5. Here, we are checking the condition whether i is greater than 100 or not, before going into the check of odd and even.

6. In the earlier example, we had given the check after the printing of the odd number. Hence even 101 was printed. 

7. This program will print numbers upto 99.

*/


