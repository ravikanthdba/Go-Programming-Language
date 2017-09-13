package main

import "fmt"

func main() {
    fmt.Println(Greet("Ravikanth ","Garimella"))
}


func Greet(first_name,last_name string) (s string){
    s = fmt.Sprint(first_name,last_name)
    return
}

/* Notes:

1. In this case, we are defining a name during the return as well.

2. This is a not so recommended approach.

*/
