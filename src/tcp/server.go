package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)

	}
}

func handleClient(conn net.Conn) {
	fmt.Println("handleClient")
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	request := make([]byte, 128)
	defer conn.Close()
	for {
		read_len, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}
		if read_len == 0 {
			fmt.Println("read_len is 0")
			break
		} else if string(request) == "timestamp" {
			fmt.Printf("request is %s", string(request))

			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			fmt.Println("else>>>")
			dayTime := time.Now().String()
			conn.Write([]byte(dayTime))
		}
		request = make([]byte, 128)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s \r\n\r\n", err.Error())
		os.Exit(1)
	}
}
