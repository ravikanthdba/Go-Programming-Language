package main

import "fmt"

func main() {
    Greet("Daniel")
    Greet("Ravikanth")
}


func Greet(first_name string) {
    fmt.Printf("Hello %s\n",first_name)
}

/* Notes:

1. In the above program, there is a function defined Greet, which is taking 1 parameter as an input which is the first name.

2. Inside the function, it captures a name, and prints "Hello <First_name>", as an output.

3. In the function main, we call the Greet function with a parameter.

4. The arguments list passed during the calling of the function is called parameters.

5. The arguments list defined during the function, is called arguments list.

6. The major use of a function is to use the same code more than once.

7. The output of the function is:

	Hello Daniel
	Hello Ravikanth

*/
