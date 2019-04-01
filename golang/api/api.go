package api

import (
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "docker"
	password = "docker"
	dbname   = "docker"
)

type Database struct {
	db *sql.DB
}

var (
	Db Database
)

func (self *Database) Open() (err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	self.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return
	}

	err = self.db.Ping()
	if err != nil {
		return
	}

	return
}

func (self *Database) Init() (err error) {
	for _, schema := range schemas {
		_, err = self.db.Exec(schema)
		if err != nil {
			return
		}
	}

	return
}

func (self *Database) Close() {
	self.db.Close()
}
