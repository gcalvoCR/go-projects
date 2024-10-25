package connection

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func (m *Postgres) Connect() error {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		"postgres",  // user
		"postgres",  // password
		"127.0.0.1", // server
		"5432",      // port
		"postgres",  // db
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	m.db = db

	fmt.Println("Connected to POSTGRES")

	return nil

}

func (m *Postgres) GetNow() (*time.Time, error) {
	var t time.Time
	err := m.db.QueryRow("select CURRENT_DATE").Scan(&t)
	if err != nil {
		fmt.Printf("error al leer la fecha del servidor: %v\n", err)
		return nil, err
	}

	return &t, nil
}

func (m *Postgres) Close() error {
	defer fmt.Println("Connection closed")
	return m.db.Close()
}
