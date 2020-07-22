package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
)

//tcp 多端口 server端

var wg sync.WaitGroup

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf []byte
	buf = make([]byte, 500)
	for {
		readLen, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("conn: %s disconnect\n", conn.RemoteAddr().String())
			} else {
				fmt.Println("conn read failed:", err)
			}
			break
		}
		fmt.Printf("conn: %s, len: %d, read data:%v\n", conn.RemoteAddr().String(), readLen, string(buf[:readLen]))
	}
}

func listenPort(port int) {
	defer wg.Done()
	address := ":" + strconv.Itoa(port)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		fmt.Printf("server port %d new conn accept\n", port)
		go process(conn)
	}
}

func main() {
	for port:=6000; port < 6010; port++ {
		wg.Add(1)
		go listenPort(port)
	}
	fmt.Println("all port listened")
	wg.Wait()
}


