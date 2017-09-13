package main

import "fmt"

func main() {
    Greet("Ravikanth","Garimella",30)
    // Greet("Ravikanth","Garimella")
    // Greet("Ravikanth","Garimella",'')
    Greet_1("Ravikanth","Garimella")
}


func Greet(first_name string,last_name string,age int) {
    fmt.Printf("Hello %s %s is of age %d\n",first_name,last_name,age)

}

func Greet_1 (first_name, last_name string) {
    fmt.Printf("Hello %s %s\n",first_name,last_name)
}

/* Notes

1. Above program shows function with multiple arguments.

2. Three arguments are passed to the function, First Name, Last Name and age.

3. The "Greet" function is called using three parameters in the first statement, which is "Ravikanth","Garimella",30.

4. In the second statement, we called only 2 parameters, which failed out.

5. In the third statement, we gave '' as the third parameter, which also had failed.

6. The difference between variable and function is, as follows:
	a. Greet() - When a variable has brackets next to it, then it will be a function.
	b. Greet   - When a variable has no brackets next to it, it will be a variable.

7. Greet_1 function arguments denote, if both the arguments are of the same type, and it can be declared once. 
	Eg : first_namee, last_name string

*/
