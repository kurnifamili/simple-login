package handler

import (
	"net"
	"encoding/json"
	"log"

	"../../common"
	dbClient "../../database/client"
)

func LoginHandler(c net.Conn) {
	d := json.NewDecoder(c)
	var req common.LoginRequest
	err := d.Decode(&req)
	if err != nil {
		log.Fatal(err)
		return
	}

	userID, _, _, err := dbClient.GetClient().GetUser(req.Username, req.Password)
	if err != nil {
		c.Write([]byte(err.Error() + "\n"))
	} else {
		c.Write([]byte("UserID " + userID + "\n"))
	}

	c.Close()
}