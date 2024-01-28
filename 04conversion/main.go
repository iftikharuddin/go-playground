package main

import (
    "bufio"
     "fmt"
      "os"
)

func main() {
    fmt.Println("Welcome to our App")
    fmt.Println("Rate it out of 5:")

    reader := bufio.NewReader(os.Stdin)

    input, _ := reader.ReadString('\n')
    fmt.Println("Thanks for rating, ", input)

    // if err != nil { fmt.Println(err) } else {}
}