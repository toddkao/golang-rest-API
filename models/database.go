package models

import (
	"log"

	re "gopkg.in/gorethink/gorethink.v4"
)

//Exported Session
var (
	Session *re.Session
	url     = "127.0.0.1"
)

// InitDB initializes connection to rethinkDB
func InitDB() {
	var err error
	Session, err = re.Connect(re.ConnectOpts{
		Address:  url,
		Database: "ecomm",
	})

	if err != nil {
		log.Fatalln(err)
	}
}
