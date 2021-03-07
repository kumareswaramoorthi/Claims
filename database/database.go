package database

import (
	"database/sql"
	"fmt"

	con "claims/constants"

	_ "github.com/lib/pq"
)

func GetConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		con.Host, con.Port, con.User, con.Password, con.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected! to db")
	return db, nil
}
