package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
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
	rSet, _ := regexp.Compile("^SET 4:([0-9\\p{L}]+) 5:([0-9\\p{L}]+)$")
	rGet, _ := regexp.Compile("^GET 4:([0-9\\p{L}]+)$")
	rDel, _ := regexp.Compile("^DEL 4:([0-9\\p{L}]+)$")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		conn.Write([]byte("Ok\n"))
		switch {
		case rSet.MatchString(line):
			fmt.Println("Setting:", line)
		case rGet.MatchString(line):
			fmt.Println("Getting:", line)
		case rDel.MatchString(line):
			fmt.Println(line, "Deleting:", line)
		case line == "exit":
			return
		}
	}
}
