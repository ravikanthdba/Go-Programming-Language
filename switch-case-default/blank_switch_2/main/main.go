package main

import "fmt"

var Rupees int = 1500

func main() {
    switch {

        case Rupees > 100 && Rupees < 400 : fmt.Printf("%d >= 100 and %d < 400\n",Rupees,Rupees)
        case Rupees >= 400 && Rupees < 900 : fmt.Printf("%d >= 400 and %d < 900\n",Rupees,Rupees)
        case Rupees >= 900 && Rupees < 1300: fmt.Printf("%d >= 900 and %d < 1300\n",Rupees,Rupees)
        default : fmt.Printf("%d < 100 or %d > 1300\n",Rupees,Rupees)
    }
}
