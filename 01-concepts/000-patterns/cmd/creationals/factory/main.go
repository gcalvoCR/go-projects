package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gcalvocr/go-patterns/cmd/creationals/factory/connection"
)

// For this example check the docker-compose file on the root folder
func main() {
	var t int
	fmt.Println("Digite 1 para conexion MYSQL y 2 para POSTGRES")
	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Printf("error al leer la opcion %v", err)
		os.Exit(1)
	}

	conn := Factory(t)

	err = conn.Connect()
	if err != nil {
		log.Fatal(err)
	}
	time, err := conn.GetNow()
	if err != nil {
		log.Println(err)
	}
	log.Printf("time: %v\n", time)
	err = conn.Close()
	if err != nil {
		fmt.Println("No se pudo consultar la fecha")
	}
}

func Factory(t int) connection.DBConnection {
	switch t {
	case 1:
		return &connection.MySQL{}
	case 2:
		return &connection.Postgres{}
	default:
		return nil
	}
}
