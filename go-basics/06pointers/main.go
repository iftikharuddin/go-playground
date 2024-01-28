package main

import "fmt"

func main() {
    fmt.Println("Welcome to Pointers")
    var ptr *int
    fmt.Println("Value of ptr:", ptr)

    myNumber := 45
    var ptr2 = &myNumber

    fmt.Println("Value of ptr:", ptr2)
    fmt.Println("Value of ptr:", *ptr2)

    *ptr2 = *ptr2 + 2
    fmt.Println("Value of ptr:", *ptr2)

}