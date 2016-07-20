package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"

	"key-value-storage-for-go-training/storage"
)

const (
	connectionType = "tcp"
)

type Cli struct {
	Host string
	Port string
}

func (c Cli) Start() {
	s := &storage.Storage{}
	s.Init()

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
		go handleCli(conn, s)
	}
}

func handleCli(conn net.Conn, s storage.SetGetDeleter) {
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
			args := rSet.FindStringSubmatch(line)
			s.Set(args[1], args[2])
			fmt.Println("Setting:", args[1], args[2])
		case rGet.MatchString(line):
			args := rGet.FindStringSubmatch(line)
			val, ok := s.Get(args[1])
			fmt.Println("Getting:", args[1])
			msg := "0: (absent)\n"
			if ok {
				msg = fmt.Sprintf("5:%s (present)\n", val)
			}
			conn.Write([]byte(msg))
		case rDel.MatchString(line):
			args := rDel.FindStringSubmatch(line)
			ok := s.Del(args[1])
			fmt.Println(line, "Deleting:", args[1], ok)
		case line == "exit":
			return
		}
	}
}
