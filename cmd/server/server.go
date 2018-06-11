package main

import (
	. "github.com/toddkao/ecommGo/internal/app/db"
	. "github.com/toddkao/ecommGo/internal/app/routes"
	// "github.com/graphql-go/graphql"
)

// main function application entry point
func main() {
	InitDB()
	InitRoutes()
}
