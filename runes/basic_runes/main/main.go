package main

import "fmt"

func main() {
    for i := 66000; i <= 66509 ; i++ {
        fmt.Printf("%v,'-',%v,'-',%v\n", i, string(i), []byte(string(i)))
    }
}

/* Notes:

1. Unique Unicode Character.

2. Every rune is a number, hence its an alias to int32. 

3. ASCII values are the initital int32 for characters.

4. Unicode is a 4 byte coding scheme.

5. In the above program, we take a value of i, we increment it from 50 to 50000, we print i, string(i) and its corresponding int32 value.

6. The runes are represented by "'" (single quotes)

*/
