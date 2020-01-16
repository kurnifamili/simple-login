package client

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"../../common"
)

var (
	client IDatabaseClient

	errInvalidCredentials = errors.New("the username and password are not valid")
)

type IDatabaseClient interface {
	GetUser(username string, password string) (string, string, string, error)
}

type DatabaseClientImpl struct {
	SQLClient *sql.DB
}

func GetClient() IDatabaseClient {
	if client == nil {
		InitClient()
	}

	return client
}

func InitClient() {
	db, err := sql.Open("mysql",
		"root@tcp(127.0.0.1:3306)/test_login")
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(common.DBMaxOpenConnections)

	client = &DatabaseClientImpl{
		SQLClient: db,
	}
}

func (m *DatabaseClientImpl) GetUser(username string, password string) (string, string, string, error) {
	stats := m.SQLClient.Stats()
	log.Printf("Current open DB conns %d\n", stats.OpenConnections)

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
