package main

import (
	// "fmt"
	// "log"
	"test-project/config"
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
)
func main() {
	db := config.GetConnection()
	defer db.Close()
	
}
