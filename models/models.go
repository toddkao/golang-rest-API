package models

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v4"
)

//Exported Session
var (
	Session *r.Session
	url     = "127.0.0.1"
)

// InitDB initializes connection to rethinkDB
func InitDB() *r.Session {
	var err error
	Session, err = r.Connect(r.ConnectOpts{
		Address:  url,
		Database: "ecomm",
	})

	if err != nil {
		log.Fatalln(err)
	}
	return Session
}
