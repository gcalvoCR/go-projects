package driver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // importing a package for its side effects (initialization)
)

func OpenDB(dns string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	// We double check the connection by pinging the DB
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
