package main

import "fmt"

func Greeter() func() string {
    return func() string{
        return "Hello World"
    }
}

func main() {
    greeting := Greeter()
    fmt.Printf("%T\n",greeting)
    fmt.Println(greeting())
}
