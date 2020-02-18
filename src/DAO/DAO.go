package DAO

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Customer struct {
	ID          int
	Name, Phone string
}

var db *sql.DB

func OpenDBConnection() error {
	database, err := sql.Open("sqlite3", "../sample.db")
	db = database
	return err
}

func CloseDBConnection() error {
	return db.Close()
}

func GetAllCustomers()  ([]Customer, error) {
	var customers []Customer
	rows, err := db.Query("select id,name,phone from customer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var phone string
		err = rows.Scan(&id, &name, &phone)

		if err != nil {
			return customers, err
		}

		customers = append(customers, Customer{id, name, phone})
	}

	return customers, nil
}
