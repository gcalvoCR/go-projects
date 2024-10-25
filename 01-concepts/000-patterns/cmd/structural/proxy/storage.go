package main

import "time"

// Data simula la base de datos
type Data struct {
	books    Books
	server   string
	port     uint16
	user     string
	password string
}

// New devuelve una nueva BD
func New(server string, port uint16, user, password string) *Data {
	d := &Data{
		server:   server,
		port:     port,
		user:     user,
		password: password,
	}
	d.load()
	return d
}

func (d *Data) ByID(ID uint) Book {
	time.Sleep(2 * time.Second)
	for _, v := range d.books {
		if v.ID == ID {
			return v
		}
	}
	return Book{}
}

func (d *Data) All() Books {
	time.Sleep(4 * time.Second)
	return d.books
}

func (d *Data) load() {
	d.books = make(Books, 0, 5)
	d.books = append(d.books,
		Book{ID: 1, Name: "El arte de la guerra", Author: "Sun Tzu"},
		Book{ID: 2, Name: "La pelota no entra por azar", Author: "Ferrian Soriano"},
		Book{ID: 3, Name: "Jesus, CEO", Author: "Laurie Beth"},
		Book{ID: 4, Name: "La biografia de Steve Jobs", Author: "Walter Issacson"},
		Book{ID: 5, Name: "Pequeño cerdo capitalista", Author: "Sofía Macías"},
	)
}
