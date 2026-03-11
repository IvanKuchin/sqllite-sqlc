package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 1. Use your exact DSN
	dsn := ":memory:?_parse_time=true" // Using in-memory for the test
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. Create a dummy table
	_, err = db.Exec("CREATE TABLE test_time (val DATETIME)")
	if err != nil {
		log.Fatal(err)
	}

	// 3. Insert using the "Clean" SQLite format
	now := time.Now().UTC()
	formatted := now.Format("2006-01-02 15:04:05")
	fmt.Printf("Inserting: %s\n", now)

	_, err = db.Exec("INSERT INTO test_time (val) VALUES (?)", formatted)
	if err != nil {
		log.Fatal(err)
	}

	// 4. Try to Scan it back into string
	var scannedString string
	err = db.QueryRow("SELECT val FROM test_time").Scan(&scannedString)
	if err != nil {
		fmt.Printf("❌ SCAN FAILED: %v\n", err)
	} else {
		fmt.Printf("✅ SCAN SUCCESS: %v (type %T)\n", scannedString, scannedString)
	}

	// 4. Try to Scan it back into time.Time
	var scannedTime time.Time
	err = db.QueryRow("SELECT val FROM test_time").Scan(&scannedTime)
	if err != nil {
		fmt.Printf("❌ SCAN FAILED: %v\n", err)
	} else {
		fmt.Printf("✅ SCAN SUCCESS: %v (type %T)\n", scannedTime, scannedTime)
	}
}
