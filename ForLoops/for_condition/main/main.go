package main

import "fmt"

func main() {
    i := 0
    for i < 10 {
        fmt.Println(i)
        i ++
    }

}

/* Notes

1. This program prints the numbers from 1 to 10.

2. This is similar to while loop in other Programming languages.

3. We do not have a while loop in Go Programming. Everything is taken care by For loop.

4. We initialize the variable first, and we keep on Printing and incrementing i until the condition becomes false.

5. It is as similar as :

while i < 10:
do
    print i
    i ++
done

6. If we forget giving an increment, the value of i will always be 0, and since 0 < 10, it will be an infinite loop.

7. After the for loop, the condition has to be a boolean value, either True or False.

*/
