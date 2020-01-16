package server

import (
	"net"
	"fmt"

	"../../common"
	"../handler"
)

func Start() {
	l, err := net.Listen("tcp", ":"+common.TcpPort)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Started TCP server at port %s..\n", common.TcpPort)
	defer l.Close()

	for {
		// accept a connection
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// handle the connection
		go handler.LoginHandler(c)
	}
}