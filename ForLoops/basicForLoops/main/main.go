package main

import "fmt"

func main() {
    for i:= 0; i <= 100; i++ {
        fmt.Println(i)
}
}

/* Notes:

1. This Program is a basic introduction to For Loops.

2. In the above Program, we print the numbers from 1 to 100.

3. The way we write a for loop is as follows:

	for initialization ; condition ; iteration {
		perform operation
	}

4. In the above program, we initialize a variable i to 0, we provide a condition that i should be less than or equal to 100, and increment i by 1 each time.

5. The following is the way it works:

	a) i is set to 0.
	b) It checks the condition, if i is less than or equal to 100.
	c) if true, it performs the operation.
	d) i will be incremented.
	e) if i satisfies the condition, then operation is performed.
	f) i will be incremented.
	g) if i satisfies the condition, then operation is performed.
	h) i is incremented.
	i) if i does not satisfy the condition, then the loop quits.

6. The other way is:
	a) i is initialized.
	b) if i does not satisfy the condition, then the loop breaks.

*/
