package main

import "fmt"

func main() {
    fmt.Println(Greet("Ravikanth ", "Garimella"))
}

func Greet(first_name, last_name string) string {
     return fmt.Sprint(first_name, last_name)
}

/* Notes

1. The above function shows the function with return type.

2. Here in the above function, we are passing first and last name as input, and we are returning a string.

3. Sprint - is the function we are using, which will print the data to the string instead of console.

4. In the function, we give the "return" statement.

5. The above program prints the name.

6. Here in the output, the space between two words is not taken automatically, and hence after "Ravikanth " (There is a small space). 

*/
