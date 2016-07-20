package main

import (
	"fmt"

	"key-value-storage-for-go-training/server"
)

func main() {
	fmt.Println("Init tcp server")
	srv := server.Cli{"localhost", "11200"}
	srv.Start()
}
