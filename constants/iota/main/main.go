package main

import "fmt"

func main() {
    const (
        a = iota;
        b = iota;
        c = iota;
    )


    const(
        x float32 = iota
        y float32 = iota
        z float32 = iota

   )

   fmt.Printf("%d\n%d\n%d\n", a,b,c)
   fmt.Printf("%f\n%f\n%f\n", x,y,z)
   fmt.Printf("%d\n%d\n%d\n", x,y,z)

}

/* Notes

1. IOTA is used for integers.

2. when defined as IOTA, the variable will be incremented automatically, until the next declaration of "CONST" is found.

3. Whenever it reaches a declaration of CONST, the value gets reset to 0

Eg: Initially there were 3 constants defined : a,b,c as IOTA. so the value will be incremented automatically.

Next there are 3 more constants defined : x,y,x. So IOTA gets reset and again starts from 0.

4. If we define the variable as a float, and while printing, we give "%d", which will be for int, we get the following output:

%!d(float32=0)
%!d(float32=1)
%!d(float32=2)

which means %d is not a float related variable.

Refer to the third Printf statement in the program.

*/
