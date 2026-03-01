package main

import (
	"database/sql"
	"log"
	"os"
)

// BAD: global variable- any function in the package can alter it
// unit tests that don't need DB are now forced to have it initialized
var db *sql.DB

// BAD: using init to open a DB connection has three problems
// 1. Error handling is limited - only option is panic, caller can't retry or fallback
// 2. init runs before All tests - even unit tests that don't need a DB connection
// 3. Forces a global variable - hard to isolate, hard to test
func init() {
	dataSourceName := os.Getenv("MYSQL_DATA_SOURCE_NAME")

	d, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err) //only option is panic - caller can't handle this
	}

	err = d.Ping()
	if err != nil {
		log.Panic(err) // same problem - forces app to stop, no retry possible
	}

	db = d // assigned to global - now every function depends on global state
}

func main() {
	//db is available but it's a global -fragile, untestable
	_ = db
}
