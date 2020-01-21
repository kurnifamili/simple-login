package handler

import (
	dbClient "../../database/client"
	"errors"
)

func loginHandler(args []string) (respStr string, respErr error){
	if len(args) != 2 {
		return "", errors.New("the request received does not follow contract" )
	}

	username := args[0]
	password := args[1]

	userID, err := dbClient.GetClient().GetUser(username, password)
	if err != nil {
		return "", err
	} else {
		return userID, nil
	}
}