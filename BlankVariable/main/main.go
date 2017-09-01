package main

import (
"fmt"
"io/ioutil"
"net/http"
)

func main() {
    res, _ := http.Get("https://golang.org/dl/")
    page, _ := ioutil.ReadAll(res.Body)
    res.Body.Close()
    fmt.Printf("%s",page)

}


/* Notes:

-    In Go, any identifier, we are declaring has to be used, or it gives an error message. While Programming, we may end up declaring some variable, but we may get some value for it or we may not get the value. Example: We write the try and catch block, where in we can receive the error, or we may not. In this case the blank identifier helps us with that. We declare it by using “-“.

*/
