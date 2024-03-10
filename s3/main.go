package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create a connection string
	connectionString := "root:1q2w3e4r5t@tcp(127.0.0.1:3306)/go_database"

	// Open a connection to the MySQL database
	db, err := sql.Open("mysql", connectionString)
	errCheck(err)

	defer db.Close()

	// Test the connection
	err = db.Ping()
	errCheck(err)

	fmt.Println("Connected to MySQL!")
	// Perform database operations...

	insert(db)

	delete(db)

	update(db)

}

func insert(db *sql.DB) {
	insert, err := db.Query("INSERT users VALUES(4,'al2')")
	insert.Close()
	fmt.Println("insert successfully!")
	errCheck(err)
}

func delete(db *sql.DB) {
	delete, err := db.Query("DELETE FROM users WHERE id=4")
	errCheck(err)
	delete.Close()
	fmt.Println("dleted!")
}

func update(db *sql.DB) {
	update, err := db.Query("UPDATE users SET name='update test' WHERE id=1")
	errCheck(err)
	update.Close()
	fmt.Println("update successfully!")
}

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}
