/*
Author:Ethan White
Version:1.0.0
Date: 25-11-19
Fileoverview: This program will take four lines and output them is reverse order
*/
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {

    // making variables
    var line1 string
    var line2 string
    var line3 string
    var line4 string

    reader := bufio.NewReader(os.Stdin)

    // inputs
    fmt.Print("Type in a line: ")
    line1, _ = reader.ReadString('\n')

    fmt.Print("Type in a line: ")
    line2, _ = reader.ReadString('\n')

    fmt.Print("Type in a line: ")
    line3, _ = reader.ReadString('\n')

    fmt.Print("Type in a line: ")
    line4, _ = reader.ReadString('\n')

    // output
    fmt.Print(line4)
    fmt.Print(line3)
    fmt.Print(line2)
    fmt.Print(line1)

    fmt.Println("\nDone.")
}
