package main

import "fmt"

func hello() {
    fmt.Println("hello ")
}

func world() {
    fmt.Println("world")
}

func main() {
    defer world()
    hello()
}

/* Notes:

1. The above program demonstrates "defer".

2. "defer" is a keyword, to hold execution of function until the end of main function.

3. In the above case, we have two functions defined : hello() amd world().

    hello() - prints the string hello
    world() - prints the string world

4. While calling the main function, we mentioned "defer" before the execution of world() function, and then called the hello() function.

5. It prints hello() first, and holds the function call having "defer" till the end of main completion.

6. Just before it completes, defer is called.

7. The above program prints:

    world
    hello

*/
