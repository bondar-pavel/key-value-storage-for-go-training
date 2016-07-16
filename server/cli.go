package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	connectionType = "tcp"
)

type Cli struct {
	Host string
	Port string
}

func (c Cli) Start() {
	hp := net.JoinHostPort(c.Host, c.Port)
	listen, err := net.Listen(connectionType, hp)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer listen.Close()
	fmt.Println("Cli server was started on", hp)
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		go handleCli(conn)
	}
}

func handleCli(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		conn.Write([]byte("Ok\n"))
		if line == "exit" {
			break
		}
	}
}
