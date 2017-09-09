package main

import "fmt"

type Contract struct {
    name string
    phone int
}

func typeofvariable(x interface{}) {
    switch x.(type) {
        case int   : fmt.Printf("The variable %v is of type %v\n",x,"int")
        case string: fmt.Printf("The variable %v is of type %v\n",x,"string")
        case Contract: fmt.Printf("The variable %v is of type %v\n",x,"Contract")
        case rune: fmt.Printf("The variable %v is of type %v\n",x,"Rune")
        default : fmt.Printf("None of the above\n")

    }
}



func main() {
    typeofvariable(7)
    typeofvariable("Ravi")
    var t = Contract{"Ravikanth",123}
    typeofvariable(t)
    typeofvariable('t')
}
