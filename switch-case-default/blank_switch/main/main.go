package main

import "fmt"

var myfriendsname = "Ravi"

func main() {
    switch {
        case len(myfriendsname) == 2 : fmt.Println("My length of my Friends name is 2.")
        case myfriendsname == "Tim": fmt.Println("Hello Tim")
        case myfriendsname == "Madhi" , myfriendsname == "Ravi" : fmt.Println("Hello Madhi or Hello Ravi")
        default : fmt.Println("He is not my friend")
    }
}


/* Notes:

1. The above program shows switch using a blank identifier.

2. In the earlier programs, we used to define the conditions in the switch and map it if true, then print 1 else print 2.

3. In this program, we define the condition at each and every case statement, and whatever matches, the corresponding operation is performed.

4. In the above case, switch has 4 conditions:

	a) If length of myfriendsname is 2, then condition 1 will be printed.
        b) If the value of myfriendsname is Tim, then "Hello Tim" will be printed.
        c) Two conditions can be specified using a ",". "," is similar to "OR". If myfriendsname is "Ravi" or "Madhi", then it prints condition 3.
        d) If none of the conditions are satisfied, it moves to default, and prints myfriendsname is not my friend.

*/
