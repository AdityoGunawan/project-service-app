package config
// go get -u github.com/go-sql-driver/mysql

import (
	"log"
	"fmt"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	dbCon := os.Getenv("db_connection") 
	db, err := sql.Open("mysql", dbCon)

	if err != nil {
		log.Fatal("error sql Open ", err.Error())
	}

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("error connect to db ", errPing.Error())
		// panic("error connect db")
	} else {
		fmt.Println("success connect to DB")
	}

	return db
	
}