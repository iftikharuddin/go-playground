package main

import (
    "github.com/ethereum/go-ethereum/ethclient"
    "fmt"
    "context"
)

var infuraURL = "https://mainnet.infura.io/v3/a35898d38da8426da726363768e1964d"

func main() {
    client, err := ethclient.DialContext(context.Background(), infuraURL)
    if err != nil {
        fmt.Println(err)
    }
    defer client.Close()
    block, err := client.BlockByNumber(context.Background(), nil)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(block.Number())
}