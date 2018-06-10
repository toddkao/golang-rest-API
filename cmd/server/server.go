package main

import (
	"fmt"
	"log"
	"net/http"

	l "github.com/toddkao/ecommGo/internal/app/locations"
	// "github.com/graphql-go/graphql"
	"github.com/gorilla/mux"
	r "gopkg.in/gorethink/gorethink.v4"
)

var url = "127.0.0.1"

// main function to boot up everything
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/locations", l.GetPeople).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	if err != nil {
		log.Fatalln(err)
	}

	res, err := r.Expr("Hello World").Run(session)
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
