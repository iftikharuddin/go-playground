package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Time study in golang")

    presentTime := time.Now()
    fmt.Println(presentTime)
    fmt.Println(presentTime.Format("01-02-2006")) // for year
    fmt.Println(presentTime.Format("01-02-2006 Monday")) // for day
    fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // time

    createdDate := time.Date(2020, time.May, 10, 23, 23, 0, 0, time.UTC)
    fmt.Println(createdDate)
    fmt.Println(createdDate.Format("01-02-2006 Monday"))

}