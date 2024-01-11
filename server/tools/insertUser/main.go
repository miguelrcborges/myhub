package main

import (
	"golang.org/x/crypto/bcrypt"

	"fmt"
	"os"

	"database/sql"
	_ "modernc.org/sqlite"
)

func main() {
	args := os.Args

	if (len(args) != 4) {
		fmt.Fprint(os.Stderr, "Usage: uins <name> <password> <db.file>")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(args[2]), 0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to hash the password: %s.\n", err.Error())
	}

	db, err := sql.Open("sqlite", args[3]);
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open database file: %s.\n", err.Error())
	}

	_, err = db.Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)", args[1], hash)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to add user to the database: %s.\n", err.Error())
	}
}
