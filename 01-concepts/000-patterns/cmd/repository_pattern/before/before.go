package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int
	Name string
}

func GetUserByID(db *sql.DB, id int) (User, error) {
	var user User
	row := db.QueryRow("SELECT id, name FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		return user, err
	}
	return user, nil
}

func mainBefore() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	user, err := GetUserByID(db, 1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User:", user)
	}
}
