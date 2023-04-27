package mocks

import (
	"database/sql"

	"github.com/DATA-DOG/go-txdb" // instalar el package
	"github.com/go-sql-driver/mysql"
)

func initDB() (db *sql.DB, err error) {

	db, err = sql.Open("txdb", "identifier")
	return
}

func init() {
	dns := mysql.Config{ //creamos el string connection dinamicamente
		User:   "root",
		Passwd: "",
		DBName: "my_db",
		Addr:   "127.0.0.1",
	}
	txdb.Register("txdb", "mysql", dns.FormatDSN())
}
