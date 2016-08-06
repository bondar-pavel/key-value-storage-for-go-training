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

func handleCli(conn net.Conn, s storage.Querier) {
	defer conn.Close()
	rSet, _ := regexp.Compile("^SET 4:([0-9\\s\\p{L}]+) 5:([0-9\\s\\p{L}]+)$")
	rGet, _ := regexp.Compile("^GET 4:([0-9\\s\\p{L}]+)$")
	rDel, _ := regexp.Compile("^DEL 4:([0-9\\s\\p{L}]+)$")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		cmd := storage.Command{}
		if rSet.MatchString(line) {
			args := rSet.FindStringSubmatch(line)
			cmd.Key, cmd.Value = args[1], args[2]
			s.Query("set", cmd)
		} else if rGet.MatchString(line) {
			args := rGet.FindStringSubmatch(line)
			cmd.Key = args[1]
			data := s.Query("get", cmd)
			msg := "0: (absent)\n"
			if data.Exists {
				msg = fmt.Sprintf("5:%s (present)\n", data.Value)
			}
			conn.Write([]byte(msg))
		} else if rDel.MatchString(line) {
			args := rDel.FindStringSubmatch(line)
			cmd.Key = args[1]
			data := s.Query("del", cmd)
			msg := "0: (absent)\n"
			if data.Exists {
				msg = "1: (deleted)\n"
			}
			conn.Write([]byte(msg))
		} else if line == "exit" {
			return
		}
	}
}
