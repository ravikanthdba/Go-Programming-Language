package main

import "fmt"

func main() {
    for init := 0; init <= 100; init ++ {
        fmt.Printf("%v - %v - %v\n",init,string(init),[]byte(string(init)))
    }

    variable := 'a'
    fmt.Printf("The value of variable is %v\n",variable)
    fmt.Printf("The type of variable is %T\n",variable)
    fmt.Printf("The type of rune(variable) is %T\n",rune(variable))
    fmt.Printf("The type of rune(rune(variable)) is %T\n",rune(rune(variable)))
    fmt.Printf("The value of rune(rune(variable)) is %v\n",rune(rune(variable)))
    fmt.Printf("The value of rune(rune(variable)) is %v\n",rune(rune(variable)))
    fmt.Printf("The value of string(rune(variable)) is %v\n",string(rune(variable)))
    fmt.Printf("The value of string(82) is %v\n",string(82))

    the_string := "Hello World"
    fmt.Printf("The value of the_string is %s\n", the_string) 
    fmt.Printf("The value of rune(the_string) is %v\n", []byte((the_string)))
}

/* Notes:

1. The above example shows more on Runes.

2. We assigned a variable "variable", to a value of 'a'

3. Since its ssigned in single quots "''", its marked as a rune.

4. The value of variable will be ASCII value of letter "a".

5. The type of variable is int32, since all runes are integers.

6. The type of rune(variable) will become rune(97), which will be int32. Since both are of integer type.

7. The type of rune(rune(variable)) will be rune(rune('a')), which will be rune(rune(97) -> rune(97) -> int32.

8. The value of rune(rune(variable)) will be rune(rune('a')) will be 97.

9. The value of string(rune(variable)) will be string(rune('a')), which will be string(97), which is "a".

10. The value of string(82), will be the string with ASCII value 82, which is "R"

11. We define a string the_string to "Hello World".

12. The string is a collection of runes.

13. We split the string in the form of bytes like []byte(the_string), it gives the list of Runes.

*/
