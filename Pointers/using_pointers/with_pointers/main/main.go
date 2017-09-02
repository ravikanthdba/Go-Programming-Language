package main

import "fmt"

func modify(z *int){
    fmt.Printf("While entering: z = %d\n", *z)
    fmt.Printf("While entering: Address of z = %d\n", z)
    *z = 0
    fmt.Printf("While exiting: z = %d\n", *z)
    fmt.Printf("While exiting: Address of z = %d\n", z)
}

func main(){
    x := 42
    fmt.Printf("Before modifying: x = %d\n",x)   
    fmt.Printf("Before modifying: Address of x = %d\n",&x)   
    modify(&x)
    fmt.Printf("After modifying: x = %d\n",x)   
    fmt.Printf("After modifying: Address of x = %d\n",&x)   

}
