package main

import (
    "github.com/ethereum/go-ethereum/crypto"
    "fmt"
    "github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
    // private key
    privateKey, err := crypto.GenerateKey() // this returns PK and err
    if err != nil {
        fmt.Println(err)
    }

    // private key convert to string
    privateData := crypto.FromECDSA(privateKey)
    fmt.Println("Private Key String: ", hexutil.Encode(privateData))

    // public key generation, we generate public key from PK via ECDSA
    publicData := crypto.FromECDSAPub(&privateKey.PublicKey)
    fmt.Println("Public Key Hash: ", hexutil.Encode(publicData))

    // now let's get address from this big public key string
    fmt.Println("Your wallet address: ", crypto.PubkeyToAddress(privateKey.PublicKey).Hex())
}