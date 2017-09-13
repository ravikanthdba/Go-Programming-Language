package main

import "fmt"

func Greet(first_name,last_name string) (string,string) {
    return fmt.Sprint(last_name,first_name) , fmt.Sprint(first_name,last_name)
//    return fmt.Sprint(first_name,last_name) , fmt.Sprint(last_name,first_name)
}


func main() {
   fmt.Println(Greet("Ravikanth ","Garimella "))
}

/* Notes:

1. We can have multiple return values.

2. Multiple returns meaning, under the same return, we can have defined number of return statements.

3. In the above example, we have 2 return values, so in the return statement, we are returning 2 values.

4. How many ever returns we give in a function, the first one which gets executed, is being displayed on the screen, and the rest of the return statements are ignored.

*/
