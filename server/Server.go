package server

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"ntrip/tcp"
)

type HandleConn struct {
}

func (handle *HandleConn) HandleConn(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Println("链接关闭", err)
			return
		}
		fmt.Print(conn.RemoteAddr().String() + ":" + msg)
	}
}

func Start(ip string, port string) {
	tcp.Listen(ip, port, &HandleConn{})
}
