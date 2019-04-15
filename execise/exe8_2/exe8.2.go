package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Errorf("%v\n", err)
		os.Exit(1)
	}
	for {
		con, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleRequest(con)
	}
}

func handleRequest(con net.Conn) {
	s := bufio.NewScanner(con)
	for s.Scan() {
		token := s.Text()
		switch token {
		case "ls":
			files, err := ioutil.ReadDir("./")
			if err != nil {
				fmt.Fprintf(con, "%v\n", err)
			} else {
				s := ""
				for i, f := range files {
					s += f.Name()
					if i != len(files) {
						s += ",\n"
					}
				}
				fmt.Fprintf(con, "%v\n", s)
			}
			break

		case "close":
			fmt.Fprintf(con, "%v\n", "closing now")
			time.Sleep(time.Second * 1)
			con.Close()
			break
		default:
			if strings.HasPrefix(token, "cd") {
				home, err := os.Getwd()
				if err != nil {
					fmt.Fprintf(con, "%v\n", err)
				}
				info := strings.Split(token, " ")
				path := home + "/" + info[1]
				fmt.Println("current directory is:" + path)
				if err := os.Chdir(path); err != nil {
					fmt.Fprintf(con, "%v\n", err)
				} else {
					home = os.Getenv("HOME")
					fmt.Fprintf(con, "enter into: %v\n", path)
				}
			} else if strings.HasPrefix(token, "get") {
				info := strings.Split(token, " ")
				data, err := ioutil.ReadFile("./" + info[1])
				if err != nil {
					fmt.Fprintf(con, "%v\n", err)
				} else {
					fmt.Fprintf(con, "read content from %s: %v\n", info[1], string(data))
				}
			}
			break
		}
	}
}
