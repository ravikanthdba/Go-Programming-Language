package main

// For any executable code, the package name needs to be given as "main". The code will become executable, only if the code has "main" function in it, else the code is only compilable and can be used in other code as the package. Please refer to the code under stringutil foldeer, for more details.

import (
    "fmt"
    "github.com/ravikanthdba/Go-Programming-Language/Packages/stringutil"
)

// When we are giving the user defined Package, we need to give the absolute path as shown above. For importing the package stringutil, we needed to give "github.com/ravikanthdba/Go-Programming-Language/Packages/stringutil".


// When we are calling the variable from a different Go Program, which caters to a package, we need to refer the variable as "package_name"."variable_name". For Example, for referencing the variable "Name", which is available in "name.go" file in the package stringutil, it needs to be referred as "stringutil.Name" as shown below.


func main() {
    var name_from_main_function string = "Ravikanth Garimella"

    fmt.Printf("Name from main Function:%s\n ",name_from_main_function)

    fmt.Printf("Name from stringutil name.go file is %s\n", stringutil.Name)

    fmt.Printf("Name from the Function GetName() is %s\n", stringutil.GetName())

    fmt.Printf("Name from the Function GetNamePassVariable after passing the variable is %s\n", stringutil.GetNamePassVariable("Ravikanth Garimella"))

    fmt.Printf("Name trying to call variable with lower case: %s\n", stringutil.GetNamePassVariable2("Nasser Hussain"))

//    fmt.Printf("Name trying to call variable with lower case directly: %s\n",stringutil.get_name_lower_case("Lower Case"))
/* Above statement fails with error : # command-line-arguments
./main.go:23: cannot refer to unexported name stringutil.get_name_lower_case */

}
