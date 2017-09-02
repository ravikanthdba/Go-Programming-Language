package main

import "fmt"

func modify(z int){
    fmt.Printf("While entering the modify function, the value of z is %d\n",z)
    fmt.Printf("While entering the address of x is %d\n",&z)
    z = 0;
    fmt.Printf("While exiting the modify function, the value of z is %d\n",z)
    fmt.Printf("While exiting the address of x is %d\n",&z)
}


func main(){
 x := 5;
 fmt.Printf("Before modifying the value of x is %d\n",x)
 fmt.Printf("Before modifying the address of x is %d\n",&x)
 modify(x)
 fmt.Printf("After modifying teh value of x is %d\n",x)
 fmt.Printf("After modifying the address of x is %d\n",&x)
}
