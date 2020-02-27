package client

import (
	"database/sql"
	"errors"
	"log"

	"../../common"
	_ "github.com/go-sql-driver/mysql"
)

var (
	client IDatabaseClient

	stmtGetUser *sql.Stmt

	errInvalidCredentials = errors.New("the username and password are not valid")
)

type IDatabaseClient interface {
	GetUser(username string, password string) (userID string, err error)
}

type databaseClientImpl struct {
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
	db.SetMaxIdleConns(common.DBMaxIdleConnections)
	//db.SetConnMaxLifetime(time.Duration(time.Second * common.DBConnMaxLifetimeInSecs))
	initStatements(db)

	client = &databaseClientImpl{
		SQLClient: db,
	}
}

func initStatements(db *sql.DB) {
	var err error
	stmtGetUser, err = db.Prepare("select id, username, password from users where username = ?")
	if err != nil {
		log.Fatal(err)
	}
}

func (m *databaseClientImpl) GetUser(username string, password string) (userID string, err error) {
	stats := m.SQLClient.Stats()
	log.Printf("open DB conns %d, idle conns %d\n", stats.OpenConnections, stats.Idle)

	row := stmtGetUser.QueryRow(username)

	var (
		respUserID   string
		respUsername string
		respPassword string
	)
	switch err := row.Scan(&respUserID, &respUsername, &respPassword); err {
	case sql.ErrNoRows:
		return "", errInvalidCredentials
	case nil:
		if respPassword != password {
			return "", errInvalidCredentials
		}
		return respUserID, nil
	default:
		log.Printf(err.Error())
		return "", err
	}
}
