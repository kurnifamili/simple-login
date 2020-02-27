package main

import (
	"database/sql"
	"log"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
)

const minStringLength = 1

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	db, err := sql.Open("mysql",
		"root@tcp(127.0.0.1:3306)/test_login")
	if err != nil {
		log.Printf(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users(username, password) VALUES(?, ?)")
	if err != nil {
		log.Printf(err.Error())
	}

	// Populate DB with 10 million users
	for i := 0; i < 10000000; i++ {
		_, err = stmt.Exec(randomString(32), randomString(32))
		if err != nil {
			fmt.Println(fmt.Sprintf("error %v", err))
			// Error is most likely caused by duplicate entry, just retry in this case
			i--
		}
	}
}

func randomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&()-{}[]?")

	b := make([]rune, rand.Intn(n-minStringLength)+minStringLength)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
