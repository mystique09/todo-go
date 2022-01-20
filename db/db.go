package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	//"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "mystique09"
	password = "mystique09"
	dbname   = "codegram"
)

func InitDb() *sql.DB {
	psqlConfig := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	conn, err := sql.Open("postgres", psqlConfig)

	if err != nil {
		fmt.Println(err)
	}
	err = conn.Ping()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Database connected!")
	return conn
}