/*-------------------------------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See https://go.microsoft.com/fwlink/?linkid=2090316 for license information.
 *-------------------------------------------------------------------------------------------------------------*/

package main
import (
    "fmt"
    "os"

)

/*
this is the main function
*/
func main(){
    fmt.Println("Hi there !")
    fmt.Println( os.Getenv("PATH") )
}