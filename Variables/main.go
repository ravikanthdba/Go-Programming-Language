package main

import "fmt"

func main() {
    /* This method is called initialization. Short Hand Notation */

     a := 32       // initialize a variable
     fmt.Println("This line is using Println: %d is of type %T", a,a) // This by default goes to the new line once the program is completed.
     fmt.Printf("This line is using Printf: %d is of type %T\n", a,a) // This doesn't create a new line once the program is completed. We need to explicitly define the \n which is for creating the new line. Please refer to 3rd print statement for more clarity.

     // This method is called declaration

     var str string   
     str = "This program describes the different types of Variable declaration"
     fmt.Println("This line is using Println command. %s is of type %T", str,str )
     fmt.Printf("This line is using Printf command. %s is of type %T\n", str,str)

    // This method of decaration of Variable is called Assign

     var decimal float32 = 1.23456789
     fmt.Println("The decimal value %f is of type %T", decimal,decimal)
     fmt.Printf("The decimal value %f is of type %T\n", decimal,decimal)


     /* Notes : When we print the statement through "Println" instead of "Printf", %d, %T, or any formatting values will not be taken into consideration. Only with "Printf", we see the %f, %T into consideration. See the execution of this program for more details. */

    /* Zero Value */

    var variable1 int;
    var variable2 float32;
    var variable3 float64;
    var variable4 bool;
    var variable5 string;
    fmt.Printf("%v is of type %T\n", variable1,variable1);
    fmt.Printf("%v is of type %T\n", variable2,variable2);
    fmt.Printf("%v is of type %T\n", variable3,variable3);
    fmt.Printf("%v is of type %T\n", variable4,variable4);
    fmt.Printf("%v is of type %T\n", variable5,variable5);

    /* Notes : If we do not declare any variable, but define them, it takes the default values.
       The default value for Integer is 0.
       The default value for Float32 is 0.
       The default value for Float64 is 0.
       The default value for boolean is false.
       The default value for string is "". */
    
    /* Zero value works only during the method of "Assign". The below throws an error message 
     variablen :=  ;
     fmt.Println("%v",variablen);
    */

    /* Boolean Values */

    /* The Logical Operators are given by 
       "&&" for "AND" 
       "||" for "OR"
       "!"  for "NOT"
       The values for boolean are assign by "true" and "false" (Everything in lower case)
    */

    var bool_val_a bool = true;
    var bool_val_b bool = false;
    var bool_val_result bool;
    bool_val_result = bool_val_a && bool_val_b;
    fmt.Printf("The result of %t and %t is %t\n",bool_val_a,bool_val_b,bool_val_result)
    bool_val_result = bool_val_a || bool_val_b;
    fmt.Printf("The result of %t or %t is %t\n",bool_val_a,bool_val_b,bool_val_result)
    bool_val_result = !(bool_val_a || bool_val_b);
    fmt.Printf("The result of !(%t or %t) is %t\n",bool_val_a,bool_val_b,bool_val_result)
    bool_val_result = !(bool_val_a && bool_val_b);
    fmt.Printf("The result of !(%t and %t) is %t\n",bool_val_a,bool_val_b,bool_val_result)

}
