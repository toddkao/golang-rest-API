package db

import (
	"fmt"
	"log"

	r "gopkg.in/gorethink/gorethink.v4"
)

var url = "127.0.0.1"
var session *r.Session

// InitDB initializes database connection
func InitDB() {
	session, err := r.Connect(r.ConnectOpts{
		Address:  url,
		Database: "ecom",
	})

	fmt.Println(session)
	if err != nil {
		log.Fatalln(err)
	}
}

// Expr runs query on session
func Expr(query string) {
	res, err := r.Expr(query).Run(session)
	if err != nil {
		log.Fatalln(err)
	}

	var response string
	err = res.One(&response)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(response)
}
