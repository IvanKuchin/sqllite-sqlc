package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	queries "test-go-sqlite/internal/database/queries"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 1. Use your exact DSN
	dsn := "./data/sqlite.db?_journal_mode=WAL&_busy_timeout=5000&_synchronous=NORMAL&_cache_size=1000000&_foreign_keys=on" // Using in-memory for the test
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	q := queries.New(db)
	ctx := context.Background()

	// // 2. Create a dummy table
	// _, err = db.Exec("CREATE TABLE test_time (val DATETIME)")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 3. Insert using the "Clean" SQLite format
	now := time.Now().UTC()
	// formatted := now.Format("2006-01-02 15:04:05")
	fmt.Printf("Inserting: %s\n", now)

	// _, err = db.Exec("INSERT INTO test_time (val) VALUES (?)", now)
	id, err := q.InsertTestTime(ctx, queries.InsertTestTimeParams{
		Val:  sql.NullTime{Time: now, Valid: true},
		Val2: now,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted ID: %d\n", id)

	// 4. Try to Scan it back into string
	var scannedString string
	err = db.QueryRow("SELECT val FROM test_time WHERE rowid = ?", id).Scan(&scannedString)
	if err != nil {
		fmt.Printf("❌ SCAN FAILED: %v\n", err)
	} else {
		fmt.Printf("✅ SCAN SUCCESS: %v (type %T)\n", scannedString, scannedString)
	}

	// 4. Try to Scan it back into time.Time
	var scannedTime time.Time
	err = db.QueryRow("SELECT val FROM test_time WHERE rowid = ?", id).Scan(&scannedTime)
	if err != nil {
		fmt.Printf("❌ SCAN FAILED: %v\n", err)
	} else {
		fmt.Printf("✅ SCAN SUCCESS: %v (type %T)\n", scannedTime, scannedTime)
	}

	testTime, err := q.GetTestTimeByID(ctx, id)
	if err != nil {
		fmt.Printf("❌ SCAN FAILED: %v\n", err)
	} else {
		fmt.Printf("✅ SCAN SUCCESS: %v (type %T)\n", testTime.Val.Time, testTime.Val.Time)
		fmt.Printf("✅ SCAN SUCCESS: %v (type %T)\n", testTime.Val2, testTime.Val2)
	}

}
