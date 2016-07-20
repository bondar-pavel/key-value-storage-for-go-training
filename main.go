package main

import (
	"fmt"

	"key-value-storage-for-go-training/storage"
	"key-value-storage-for-go-training/server"
)

func main() {
	s := storage.Storage{}
	s.Init()
	key := "key1"
	value := "value1"
	fmt.Printf("Setting key '%s', value '%s'\n", key, value)
	s.Set(key, value)
	val, _ := s.Get(key)
	fmt.Printf("Getting key '%s', received '%s'\n", key, val)
	ok := s.Del(key)
	fmt.Printf("Deleting key '%s', found: %s'\n", key, ok)

	fmt.Println("Init tcp server")
	srv := server.Cli{"localhost", "11200"}
	srv.Start()
}
