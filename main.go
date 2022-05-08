package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var (
		err  error
		ln   net.Listener
		conn net.Conn
	)

	if ln, err = net.Listen("tcp", ":8080"); err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}
	for {
		if conn, err = ln.Accept(); err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		}
		go process(conn)
	}
}
func process(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		format := time.Now().Format("2006-01-02 15:04:05")
		if err != nil {
			fmt.Printf("time=%s read from connect failed, err: %v\n", format, err)
			break
		}
		fmt.Printf("time=%s receive from client, data: %v\n", format, string(buf[:n]))
		if _, err = conn.Write([]byte("Send From Server")); err != nil {
			fmt.Printf("time=%s write to client failed, err: %v\n", format, err)
			break
		}
	}
}
