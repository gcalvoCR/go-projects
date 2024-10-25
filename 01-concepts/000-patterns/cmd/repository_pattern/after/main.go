package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Model: The User struct stays the same.
// Repository Interface: Defines the methods to interact with the data layer.
// Repository Implementation: Implements the data access logic.
// Main: Uses the repository to fetch data, abstracting away the details of data access.

// User model
type User struct {
	ID   int
	Name string
}

// UserRepository defines the methods to be implemented by any data source
type IUserRepository interface {
	GetUserByID(id int) (User, error)
}

// MySQLUserRepository implements UserRepository for MySQL database
type MySQLUserRepository struct {
	db *sql.DB
}

// NewMySQLUserRepository creates a new MySQLUserRepository
func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

// GetUserByID retrieves a user by ID from MySQL
func (repo *MySQLUserRepository) GetUserByID(id int) (User, error) {
	var user User
	row := repo.db.QueryRow("SELECT id, name FROM users WHERE id = ?", id)
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		return user, err
	}
	return user, nil
}

// Business logic layer (or service layer) which is decoupled from data access
func GetUserProfile(repo IUserRepository, id int) {
	user, err := repo.GetUserByID(id)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User:", user)
	}
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Use the repository to interact with the database
	userRepo := NewMySQLUserRepository(db)
	GetUserProfile(userRepo, 1)
}
