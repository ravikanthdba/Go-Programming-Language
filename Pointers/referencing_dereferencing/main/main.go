package main

import "fmt"

func main(){
    x := 42;
    fmt.Printf("The value of x is %d\n",x)
    var y *int = &x;
    fmt.Printf("The value of y is %v\n",y)
    fmt.Printf("The memory address of y is %v\n",&y)
    fmt.Printf("The value of *y is %v\n",*y)
    fmt.Printf("The value of &*y is %v\n",&*y)
//    fmt.Printf("The value of **y is %v",**y)
//    fmt.Printf("The value of &&y is ",&&y)
//    fmt.Printf("The value of &(*y) is ",&(*y))
//    fmt.Printf("The value of *(&y) is ",*(&y))

}
