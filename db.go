package receipt

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db = sqlx.MustConnect("sqlite3", "./db.sqlite3")

func init() {
	db.MustExec(`CREATE TABLE IF NOT EXISTS product (
		id INTEGER PRIMARY KEY,
		price FLOAT,
		name VARCHAR(128),
		unitname VARCHAR(16)
    );`)
	db.MustExec(`CREATE TABLE IF NOT EXISTS discount (
		id INTEGER PRIMARY KEY,
		discounttype INTEGER,
		description VARCHAR(128) NULL,
		disabled TEXT NULL,
		productIds TEXT NULL
	);`)

}
