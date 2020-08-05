package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	//importing pg driver
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("postgres", "host=192.168.99.100 port=54320 user=postgres dbname=users password=Welcome@123 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	err = CreateSchema()
	if err != nil {
		log.Fatalln(err)
	}
}

//CreateSchema function to setup db schema
func CreateSchema() error {
	insertQuery := `CREATE TABLE IF NOT EXISTS loggedinuser (id int, name text, email text UNIQUE, meta jsonb)`
	_, err := db.Exec(insertQuery)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
