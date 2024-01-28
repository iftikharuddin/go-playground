package main

import (
    "bufio"
     "fmt"
      "os"
)

func main() {
    welcome := "Welcome"
    fmt.Println(welcome)

    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter your input:")

    // comma ok || err err
    input, _ := reader.ReadString('\n')
    fmt.Println("Thanks for the input:", input)
    fmt.Printf("Thanks for the input %T:", input)

}