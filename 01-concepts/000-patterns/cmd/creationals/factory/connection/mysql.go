package connection

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	db *sql.DB
}

func (m *MySQL) Connect() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true&allowNativePasswords=true&parseTime=true",
		"root",
		"root",
		"127.0.0.1",
		"3306",
		"mysql",
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	m.db = db

	fmt.Println("Connected to MYSQL")

	return nil

}

func (m *MySQL) GetNow() (*time.Time, error) {
	var t time.Time
	err := m.db.QueryRow("select CURDATE()").Scan(&t)
	if err != nil {
		fmt.Printf("error al leer la fecha del servidor: %v\n", err)
		return nil, err
	}

	return &t, nil
}

func (m *MySQL) Close() error {
	defer fmt.Println("Connection closed")
	return m.db.Close()
}
