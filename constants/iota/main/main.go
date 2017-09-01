package main

import "fmt"

func main() {
    const (
        a = iota;
        b = iota;
        c = iota;
    )


    const(
        x float32 = iota
        y float32 = iota
        z float32 = iota

   )

   fmt.Printf("%d\n%d\n%d\n", a,b,c)
   fmt.Printf("%f\n%f\n%f\n", x,y,z)
   fmt.Printf("%d\n%d\n%d\n", x,y,z)

}
