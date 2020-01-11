package database

import (
	"database/sql"
	"log"
	"errors"
)

var (
	errInvalidCredentials = errors.New("the username and password are not valid")
)

type IDatabaseClient interface {
	GetUser(username string, password string) (string, string, string, error)
}

type DatabaseClientImpl struct {
	SQLClient *sql.DB
}

func NewDatabaseClient() IDatabaseClient {
	db, err := sql.Open("mysql",
		"root@tcp(127.0.0.1:3306)/test_login")
	if err != nil {
		log.Fatal(err)
	}

	return &DatabaseClientImpl{
		SQLClient: db,
	}
}

func (m *DatabaseClientImpl) GetUser(username string, password string) (string, string, string, error) {
	row := m.SQLClient.QueryRow("select id, username, password from users where username = ?", username)

	var (
		respUserID string
		respUsername string
		respPassword string
	)
	switch err := row.Scan(&respUserID, &respUsername, &respPassword); err {
	case sql.ErrNoRows:
		return "", "", "", errInvalidCredentials
	case nil:
		if respPassword != password {
			return "", "", "", errInvalidCredentials
		}
		return respUserID, respUsername, respPassword, nil
	default:
		panic(err)
		return "", "", "", err
	}
}
