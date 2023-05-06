package tcp

import (
	"fmt"
	"net"
)

type Handle interface {
	HandleConn(*net.TCPConn)
}

func Listen(ip string, port string, handle Handle) {
	addr := ip + ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	if err != nil {
		panic("解析ip地址失败: " + err.Error())
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic("监听TCP失败: " + err.Error())
	}
	fmt.Println("Listen success on" + addr + "with tcp4")

	defer func() {
		fmt.Println("Close listenning ....")
		listener.Close()
		fmt.Println("Shutdown")
	}()

	go accept(listener, handle)

	select {}
}

func accept(listener *net.TCPListener, handle Handle) {
	for {
		connection, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept 失败: " + err.Error())
		} else {
			go handle.HandleConn(connection)
		}
	}
}
