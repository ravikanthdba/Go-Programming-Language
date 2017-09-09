package main

import "fmt"

func main() {
    the_string := "Hello World"
    fmt.Printf("The value of the string is %v\n",the_string)
    fmt.Printf("The type  of the string is %T\n",the_string)
    fmt.Printf("The runes list of the string %v is %v\n",the_string,[]byte(the_string))
    fmt.Printf("The runes list of the string %v is %T\n",the_string,[]byte(the_string))

    for number := 0; number <= 500 ; number++ {
        if number == 72 || number == 101 || number == 108 || number == 111 {
           fmt.Printf("%v - %v\n",number, string(number))
       }
    }
}
