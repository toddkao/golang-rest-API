package models

import (
	"log"

	re "gopkg.in/gorethink/gorethink.v4"
)

//Session exported
var (
	Session *re.Session
	url     = "127.0.0.1"
)

// Init initializes connection to rethinkDB
func Init() {
	var err error
	Session, err = re.Connect(re.ConnectOpts{
		Address:  url,
		Database: "ecomm",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
