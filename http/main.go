package main

import "github.com/hschain/go-sdk/http/server"

func main() {
	server := server.NewServer("0.0.0.0:26657", "https://scan.hschain.io/api/v2")
	//server := server.NewServer("0.0.0.0:26657", "https://testnet.hschain.io/api/lcd")
	server.Router()
}
