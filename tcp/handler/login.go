package handler

import (
	"net"
	"encoding/json"
	"log"

	"../../common"
	dbClient "../../database/client"
	"io"
)

func LoginHandler(c net.Conn) {
	for {
		//netData, err :=bufio.NewReader(c).ReadString('\n')
		d := json.NewDecoder(c)
		var req common.LoginRequest
		err := d.Decode(&req)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("error in decoding request", err)
			continue
		}

		userID, _, _, err := dbClient.GetClient().GetUser(req.Username, req.Password)
		if err != nil {
			c.Write([]byte(err.Error() + "\n"))
		} else {
			c.Write([]byte("UserID " + userID + "\n"))
		}
	}

	c.Close()
}