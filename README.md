# go-playground

## For Go Basics
( https://youtu.be/JoJ8Sw5Yb4c?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa )

## Golang-ethereum
( https://www.youtube.com/playlist?list=PLay9kDOVd_x7hbhssw4pTKZHzzc6OG0e_ )

- 01 Interact with Ethereum blockchain using Golang
- 02 Check the balance of an ether wallet
- 03 Generate an Ether Wallet Using Golang
- 04 Generate Ethereum Keystore Wallet using Golang
- 05 How to get Ethers for testing
- 06 Transferring Ether between accounts using Golang
- 07 Project Overview Smart Contract Using Solidity and Golang.
- 08 Build a To-do smart contract using Solidity.
- 09 Authorization in Ethereum Smart Contract using Solidity.
- 10-install sqlc and abigen ethereum.


## Basics learning
- Go, used by most new startups and `FAANG` are loving it
- Ideal for cloud infrastructure
- Golang is compiled
    - Gotool can run file directly without VM
    - Executables are different for OS
- System apps to web apps can be build - Cloud  ( mostly used in google products )
- Already in production
- Next Gen Language, entirely diff so don't bring other lang baggage
- Object Oriented? Yes and No, no classes but there are structs
    - no overrides in go lang
- What is missing in Go? Try/catch, but in reality you don't need it. Lexer does a lot of work
    - no semicolons, even docs say use it :P but lexer does all the work
- Golang installation and hello world
- go mod init is kinda npm init
- Go is super case sensitive
    - go env gopath != go env GOPATH
- The formal grammer use semicolons ";"
- Lexer, the job of lexer is to monitor the grammar
- in Go everything is a type
    - string, bool, uint64, complex, float32 etc
- arrays, pointers, structs, slices, mappings, and channels
- Variable should be known in advance
- Variables, types and constants
- := is not allowed outside main, you will have to use var for that
- Comma ok syntax and packages in golang
- Conversions in golang
- Handling time in golang
    - Time package
    - presentTime := time.Now()
    ````
      fmt.Println(presentTime)
      fmt.Println(presentTime.Format("01-02-2006")) // for year
      fmt.Println(presentTime.Format("01-02-2006 Monday")) // for day
      fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // time
      
                                 // month, day, hour, minute, seconds, nanoseconds, location   
      createdDate := time.Date(2020, time.May, 10, 23, 23, 0, 0, time.UTC)
      fmt.Println(createdDate)
      fmt.Println(createdDate.Format("01-02-2006 Monday"))
    `````
- Build for windows, linux and mac
    - e.g if you wanna build the app for windows ( GOOS="windows" go build)   
    - you can build standalone program in GO
- Memory management in golang
    - Memory allocation and de-allocation happens automatically 
    - new() and make()
    - new allocate memory and not init while make allocate and init the memory
    - new() you will get a memory address and zeroed storage
    - make() you will get a memory address and non-zeroed storage
    - GC(Garbage Collection) happen auto
    - NumCPU() how many CPU in used
- Pointers in golang
    - a direct reference to memory address