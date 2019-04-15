package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}
func handleConn(c net.Conn) {
	tick := time.NewTicker(1 * time.Second)
	enter := make(chan string)
	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			enter <- input.Text()
		}
		fmt.Println("Closing")
	}()

	for countdown := 10; countdown >= 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick.C:
		case str := <-enter:
			go echo(c, str, 1*time.Second)
			countdown = 10
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
	close(enter)
	tick.Stop()
}
