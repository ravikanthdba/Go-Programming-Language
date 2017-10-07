package main

import "fmt"

func farenheit2celcius(farenheit float64) float64 {
    var celcius float64
    celcius = (farenheit - 32) * 5 / 9
    return celcius
}


func celcius2farenheit(celcius float64) float64 {
    var farenheit float64
    farenheit = ((celcius * 9) / 5) + 32
    return farenheit
}

func main() {
    var farenheit float64
    var celcius float64
    var input_type string
    fmt.Println("Enter the type of input:")
    fmt.Scan(&input_type)
    if (input_type == "F") {
        fmt.Println("Enter the value of Farenheit:")
        fmt.Scan(&farenheit)
        fmt.Printf("The Celcius value of Farenheit value %.2fF is %.2fC\n",farenheit,farenheit2celcius(farenheit))
    } else if (input_type == "C") {
        fmt.Println("Enter the value of Celcius: ")
        fmt.Scan(&celcius)
        fmt.Printf("The Farenheit value of Celcius value %.2fC is %.2fF\n",celcius,celcius2farenheit(celcius))
    } else {
        fmt.Println("Invalid Input")
    }
}
