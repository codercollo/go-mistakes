package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// FIX: plain function instead of init
// Benefits:
// 1. Returns error - caller decides what to do (retry, fallback, panic)
// 2. Does not run automatically  -unit tests are not affected
// 3. No global variable - db is encapsulated, passed where needed
func createClient(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err //caller handles this - not our decision to panic
	}
	if err = db.Ping(); err != nil {
		return nil, err //caller handles this - counld retry, use fallback ,  etc
	}
	return db, nil //encapsulated - no global state
}

func main() {
	dsn := os.Getenv("MYSQL_DATA_SOURSE_NAME")

	//Caller is in control of error handling
	db, err := createClient(dsn)
	if err != nil {
		//caller decides = in main we choose to fatal
		//but in a library or test, caller could do something else
		log.Fatal("failed to connect to database:", err)
	}
	defer db.Close()

	fmt.Println("connected to database successfully!")
	_ = db
}
