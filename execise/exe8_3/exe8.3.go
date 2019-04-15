package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}
	con := conn.(*net.TCPConn)
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, con)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(con, os.Stdin)
	con.CloseWrite()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
