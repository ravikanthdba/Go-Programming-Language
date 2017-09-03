package main

import "fmt"

func main(){
    i := 0
    for {
        fmt.Println(i)
        i ++
        if i > 10 {
            break
        }
    }

}

/* Notes:

1. The above program shows how to use a "break" clause.

2. We initialize a variable i to 0.

3. We start the for loop without giving any condition.

4. We print the value of i, and then we increment.

5. We then mention the condition, if i > 10, then break the loop.

6. The steps of execution are:

	a) i is initialized to 0.
	b) for loop starts
	c) prints i
	d) increments i
	e) checks if i > 10 for the next run.
	f) if i < 10, then performs steps c,d,e.
	g) if i > 10, then exit out of loop.
	h) Program ends.

*/
