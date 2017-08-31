package vis

import "fmt"

func Print_function() {
    fmt.Printf("====== The output is from Print_function() function of the package vis =====")
    fmt.Printf("My name is %s\n",Myname)
    fmt.Printf("My current status is:\n%s", current_status)
}


/* Since Print_function function is inside the package "vis", we are able to access the variable current_status, and since the function Print_function starts with the Upper case, the function can be called outside the package as well. Because of this the variable "current_status" can be indirectly used outside the package, in this fashion */
