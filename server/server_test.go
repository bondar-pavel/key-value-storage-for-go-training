package server

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"key-value-storage-for-go-training/storage"
)

var userInput string
var reply string

type FakeConn struct{}

func (f FakeConn) Read(p []byte) (n int, err error) {
	buf := []byte(userInput)
	n = len(buf)
	if n > len(p) {
		n = len(p)
	}
	for i := 0; i < n; i++ {
		p[i] = buf[i]
	}
	userInput = ""
	return
}

func (f FakeConn) Write(p []byte) (n int, err error) {
	reply = reply + string(p)
	return len(p), nil
}

func (f FakeConn) Close() error {
	return nil
}

func runHandleCli(t *testing.T, input, output string) {
	userInput = input
	reply = ""
	conn := FakeConn{}
	s := &storage.Storage{}
	s.Init()
	handleCli(conn, s)
	assert.Equal(t, output, reply)
}

func TestHandleCliSet(t *testing.T) {
	input := "SET 4:sad 5:qwe\n"
	output := ""
	runHandleCli(t, input, output)
}

func TestHandleCliGet(t *testing.T) {
	input := "SET 4:sad 5:qwe\nGET 4:sad"
	output := "5:qwe (present)\n"
	runHandleCli(t, input, output)
}

func TestHandleCliGetNoValue(t *testing.T) {
	input := "GET 4:sad\n"
	output := "0: (absent)\n"
	runHandleCli(t, input, output)
}

func TestHandleCliGetUnicode(t *testing.T) {
	input := "SET 4:ключ 5:значение\nGET 4:ключ\n"
	output := "5:значение (present)\n"
	runHandleCli(t, input, output)
}

func TestHandleCliGetUnicodeSpace(t *testing.T) {
	input := "SET 4:а и б 5:в и г\nGET 4:а и б\n"
	output := "5:в и г (present)\n"
	runHandleCli(t, input, output)
}

func TestHandleCliDel(t *testing.T) {
	input := "SET 4:k 5:v\nDEL 4:k\n"
	output := "1: (deleted)\n"
	runHandleCli(t, input, output)
}

func TestHandleCliDelNotPresent(t *testing.T) {
	input := "DEL 4:k\n"
	output := "0: (absent)\n"
	runHandleCli(t, input, output)
}
