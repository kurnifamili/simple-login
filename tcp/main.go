package main

import (
	"fmt"
	"net"
	"../http/database"
	"../http/common"
	"encoding/json"
	"log"
)

const TcpPort = "8091"

var (
	dbClient = database.NewDatabaseClient()
)

func main() {
	l, err := net.Listen("tcp", ":"+TcpPort)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Started TCP server at port %s..\n", TcpPort)
	defer l.Close()

	for {
		// accept a connection
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleLogin(c)
	}
}

func handleLogin(c net.Conn) {
	d := json.NewDecoder(c)
	var req common.LoginRequest
	err := d.Decode(&req)
	if err != nil {
		log.Fatal(err)
		return
	}

	userID, _, _, err := dbClient.GetUser(req.Username, req.Password)
	if err != nil {
		c.Write([]byte(err.Error() + "\n"))
		fmt.Println("Error in getting user, err=", err.Error())
	} else {
		c.Write([]byte("UserID " + userID + "\n"))
		fmt.Println("UserID:", userID)
	}

	c.Close()
}