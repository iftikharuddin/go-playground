package main

import (
    "fmt"
    "github.com/ethereum/go-ethereum/accounts/keystore"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "io/ioutil"
)

func main() {
//     // key store
//     key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
//     // password to encrypt PK
    password := "thePassword"
//     // create new account encrypted with password
//     a, err := key.NewAccount(password)
//     if err != nil {
//         fmt.Println(err)
//     }
//     fmt.Println(a.Address)

    // now let's read the PK from the encrypted PK
    b, err := ioutil.ReadFile("./wallet/UTC--2024-01-28T19-12-19.336010000Z--befded21962a0cb681c9ab7dd7b7494ed264a1bd")
    if err != nil {
        fmt.Println(err)
    }

    // lets decrypt the encrypted PK
    key, err := keystore.DecryptKey(b, password)
    if err != nil {
        fmt.Println(err)
    }

    // convert PK to String
    privateData := crypto.FromECDSA(key.PrivateKey)
    fmt.Println(hexutil.Encode(privateData))

    // now lets get public key
    // public key generation, we generate public key from PK via ECDSA
    publicData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
    fmt.Println(hexutil.Encode(publicData))

    // now let's get address from this big public key string
    fmt.Println("Your wallet address: ", crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())

}