package main

import "fmt"

func main(){
    for i:=0 ; i < 5 ; i++ {
        for j:=0 ; j < 5 ; j++ {
            fmt.Println(i,"-",j)
        }
    }
}

/* Notes:

1. Loop inside a Loop is called nested loop.

2. In this Program, we initialize the variable of i to 0, and j to 0.

3. For each and every value of i, the inner loop will be executed 5 times, until the condition for inner loop fails.

4. Once the inner loop condition fails, it comes to the outer loop and increments the value of outer loop.

5. Until the outer loop condition fails, the inner loop is run those many number of times.

The Eg o/p in this Program is 

0 - 0
0 - 1
0 - 2
0 - 3
0 - 4

The condition of inner loop has failed, so inner loop exits, and now the outer loop i is incremented. The output becomes:

1 - 0
1 - 1
1 - 2
1 - 3
1 - 4
etc

*/


