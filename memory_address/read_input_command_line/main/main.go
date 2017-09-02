package main

import "fmt"

var pi float32 = 3.1417

func main(){
    var radius float32;
    fmt.Print("Enter the radius of the circle:");
    fmt.Scan(&radius);
    var area float32;
    area = pi * (radius * radius);
    fmt.Printf("The area of the circle with radius %f is %f\n",radius,area)
}


/* Notes

1. The above program is written to find the area of a circle, given its radius.

2. We are giving the value through the Standard Input.

3. Standard input is given through the function "Scan".

4. We are referencing to the memory address of the variable.

5. We give the value through standard input, which is linked with the radius variable declared through the memory address.

6. We declare the variable area, and define area as pi*(radius**2).

7. Since we cannot multiply a float variable with the integer, pi * (r ** 2) could ot work here as the 2 is an integer. Hence we had given all variables to be float32 and given the formula to be pi * (radius * radius)

8. Researching on how to change the type of the variable .

*/
